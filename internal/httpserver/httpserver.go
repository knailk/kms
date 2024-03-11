// Package httpserver provides a preconfigured HTTP server.
package httpserver

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"kms/app/config"
	"kms/app/errs"
	mdlwCORS "kms/app/middleware/cors"
	mdlwRecovery "kms/app/middleware/recovery"
	mdlwSecHdr "kms/app/middleware/secureheaders"
	"kms/app/service"
	"kms/internal/httpserver/driver"
	"kms/internal/shutdown"
	"kms/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

const (
	apiHealthPath = "/"
	pathPrefix    = "/api"
)

// Services are used by the application service handlers
type Services struct {
	OrgServicer            service.OrgServicer
	AppServicer            service.AppServicer
	RegisterUserService    service.RegisterUserServicer
	LoggerService          service.LoggerServicer
	GenesisServicer        service.GenesisServicer
	AuthenticationServicer service.AuthenticationServicer
	AuthorizationServicer  service.AuthorizationServicer
	PermissionServicer     service.PermissionServicer
	RoleServicer           service.RoleServicer
	MovieServicer          service.MovieServicer
}

// Server represents an HTTP server.
type Server struct {
	router *gin.Engine
	Driver driver.Server

	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string

	// Services used by the various HTTP routes and middleware.
	Services
}

type task struct {
	srv *http.Server
}

func (t *task) Name() string {
	return "httpserver"
}

func (t *task) Shutdown(ctx context.Context) error {
	if err := t.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}
	return nil
}

// New initializes a new Server and registers
// routes to the given router
func New(engine *gin.Engine, serverDriver driver.Server, tasks *shutdown.Tasks) *Server {
	if tasks.HasStopSignal() {
		return nil
	}
	s := &Server{router: engine}
	s.Driver = serverDriver

	// register routes to the router
	s.registerRoutes()

	return s
}

// ListenAndServe is a wrapper to use wherever http.ListenAndServe is used.
func (s *Server) ListenAndServe() error {
	const op errs.Op = "server/Server.ListenAndServe"
	if s.Addr == "" {
		return errs.E(op, errs.Internal, "Server Addr is empty")
	}
	if s.router == nil {
		return errs.E(op, errs.Internal, "Server router is nil")
	}
	if s.Driver == nil {
		return errs.E(op, errs.Internal, "Server driver is nil")
	}
	return s.Driver.ListenAndServe(s.Addr, s.router)
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.Driver.Shutdown(ctx)
}

// Driver implements the driver.Server interface. The zero value is a valid http.Server.
type Driver struct {
	Server http.Server
}

// NewDriver creates a Driver enfolding a http.Server with default timeouts.
func NewDriver() *Driver {
	return &Driver{
		Server: http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

// ListenAndServe sets the address and handler on Driver's http.Server,
// then calls ListenAndServe on it.
func (d *Driver) ListenAndServe(addr string, h http.Handler) error {
	d.Server.Addr = addr
	d.Server.Handler = h
	return d.Server.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any active connections,
// by calling Shutdown on Driver's http.Server
func (d *Driver) Shutdown(ctx context.Context) error {
	return d.Server.Shutdown(ctx)
}

// NewGinRouter initializes a gin-gonic/gin router
func NewGinRouter(ctx context.Context, cfg *config.Config) (*gin.Engine, error) {
	const op errs.Op = "httpserver/NewGinRouter"
	// Enable Release mode for production
	if cfg.Env == config.EnvProduction || cfg.Env == config.EnvStaging {
		gin.SetMode(gin.ReleaseMode)
	}

	// initializer gin-gonic/gin router
	router := gin.New()

	initMiddlewares(router, cfg)

	// Init routes
	if err := initRoutes(ctx, router, cfg); err != nil {
		return nil, errs.E(op, err)
	}

	return router, nil
}

// decoderErr is a convenience function to handle errors returned by
// json.NewDecoder(r.Body).Decode(&data) and return the appropriate
// error response
func decoderErr(err error) error {
	const op errs.Op = "server/decoderErr"

	switch {
	// If the request body is empty (io.EOF)
	// return an error
	case err == io.EOF:
		return errs.E(op, errs.InvalidRequest, "request body cannot be empty")
	// If the request body has malformed JSON (io.ErrUnexpectedEOF)
	// return an error
	case err == io.ErrUnexpectedEOF:
		return errs.E(op, errs.InvalidRequest, "malformed JSON")
	// return other errors
	case err != nil:
		return errs.E(op, err)
	}
	return nil
}

func initMiddlewares(router *gin.Engine, cfg *config.Config) {
	var middlewares []gin.HandlerFunc

	// Recovery middleware (recover from panic)
	middlewares = append(middlewares, mdlwRecovery.Recovery())

	// Logger middleware that skips health check endpoint
	middlewares = append(middlewares, gin.LoggerWithFormatter(logger.HTTPLogger))

	// Tracing middleware that enables tracing of requests
	if cfg.Env == config.EnvProduction || cfg.Env == config.EnvStaging {
		middlewares = append(middlewares, gintrace.Middleware(cfg.App.Name))
	}

	// Secure headers middleware
	if cfg.Env == config.EnvProduction || cfg.Env == config.EnvStaging {
		middlewares = append(middlewares, mdlwSecHdr.SecureHeaders)
	}

	// CORS middleware
	middlewares = append(middlewares, mdlwCORS.CORS(cfg))

	router.Use(middlewares...)
}

func initRoutes(
	ctx context.Context,
	router *gin.Engine,
	cfg *config.Config,
) error {
	// Base routes
	router.GET(apiHealthPath, healthz)

	// route not found
	router.NoRoute(func(ctx *gin.Context) {
		logger.Error(errs.RouteNotFound.String(), logrus.Fields{})
		errs.HTTPErrorResponse(ctx, errs.E(errs.RouteNotFound))
	})

	return nil
}

// healthz for checking service status
func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
