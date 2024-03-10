// Package httpserver provides a preconfigured HTTP server.
package httpserver

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"kms/app/errs"
	"kms/app/service"
	"kms/internal/httpserver/driver"
	"kms/internal/shutdown"
)

const pathPrefix string = "/api"

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

	// all logging is done with a zerolog.Logger
	Logger zerolog.Logger

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

// New initializes a new Server and registers
// routes to the given router

func New(engine *gin.Engine, serverDriver driver.Server, tasks *shutdown.Tasks, lgr zerolog.Logger) *Server {
	if tasks.HasStopSignal() {
		return nil
	}
	t := &task{}
	s := &Server{router: engine}
	s.Logger = lgr
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
	t := &task{}
	t.srv = &Driver{
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
func NewGinRouter() *gin.Engine {
	// Enable Release mode for production
	if cfg.Env == config.EnvProduction || cfg.Env == config.EnvStaging {
		gin.SetMode(gin.ReleaseMode)
	}

	// initializer gin-gonic/gin router
	r := gin.New()

	initMiddlewares(router, cfg)

	// Init routes
	if err := initRoutes(ctx, router, cfg, provider); err != nil {
		return nil, err
	}

	return r
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

func (t *task) Name() string {
	return "httpserver"
}

func (t *task) Shutdown(ctx context.Context) error {
	if err := t.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}
	return nil
}
