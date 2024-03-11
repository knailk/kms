// Package httpserver provides a preconfigured HTTP server.
package httpserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"kms/app/config"
	"kms/app/errs"
	mdlwCORS "kms/app/middleware/cors"
	mdlwRecovery "kms/app/middleware/recovery"
	mdlwSecHdr "kms/app/middleware/secureheaders"
	"kms/app/service"
	"kms/internal/shutdown"
	"kms/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

const (
	apiHealthPath = "/"
	apiV1         = "/api/v1"
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
func New(ctx context.Context, cfg *config.Config, tasks *shutdown.Tasks) error {
	if tasks.HasStopSignal() {
		return nil
	}

	engine := NewGinRouter(ctx, cfg)

	t := &task{
		srv: &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.HTTPServer.Port),
			Handler:           engine,
			ReadHeaderTimeout: 10 * time.Second,
		},
	}

	go func() {
		if err := t.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error(), logrus.Fields{"Message": "HTTP server ListenAndServe failed"})
			tasks.GetSigChan() <- os.Interrupt
		}
	}()

	// Add shutdown task
	tasks.Add(t)
	return nil
}

// NewGinRouter initializes a gin-gonic/gin router
func NewGinRouter(ctx context.Context, cfg *config.Config) *gin.Engine {
	const op errs.Op = "httpserver/NewGinRouter"
	// Enable Release mode for production
	if cfg.Env == config.EnvProduction || cfg.Env == config.EnvStaging {
		gin.SetMode(gin.ReleaseMode)
	}

	// initializer gin-gonic/gin router
	router := gin.New()

	initMiddlewares(router, cfg)

	// Init routes
	initRoutes(ctx, router, cfg)

	return router
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
) {
	// Base routes
	router.GET(apiHealthPath, healthz)

	// route not found
	router.NoRoute(func(ctx *gin.Context) {
		logger.Error(errs.RouteNotFound.String(), logrus.Fields{})
		errs.HTTPErrorResponse(ctx, errs.E(errs.RouteNotFound))
	})

	return
}

// healthz for checking service status
func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
