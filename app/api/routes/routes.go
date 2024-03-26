package routes

import (
	"context"
	"kms/app/config"
	"kms/app/errs"
	mdlwCORS "kms/app/middleware/cors"
	mdlwRecovery "kms/app/middleware/recovery"
	mdlwSecHdr "kms/app/middleware/secureheaders"
	"kms/app/registry"
	"kms/internal/shutdown"
	"kms/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

const (
	apiHealthPath = "/"
	apiCommonV1   = "/api/v1"
	apiStudentV1  = "/api/v1/student"
	apiTeacherV1  = "/api/v1/teacher"
	apiDriverV1   = "/api/v1/driver"
	apiChefV1     = "/api/v1/chef"
	// apiAdminV1     = "/api/v1/admin"
)

// NewGinRouter initializes a gin-gonic/gin router
func NewGinRouter(ctx context.Context, cfg *config.Config, tasks *shutdown.Tasks, provider *registry.Provider) *gin.Engine {
	// Enable Release mode for production
	if cfg.Env == config.EnvProduction || cfg.Env == config.EnvStaging {
		gin.SetMode(gin.ReleaseMode)
	}

	// initializer gin-gonic/gin router
	router := gin.New()

	initMiddlewares(router, cfg)

	// Init routes
	initRoutes(ctx, router, cfg, provider)

	return router
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
	provider *registry.Provider,
) {
	// Base routes
	router.GET(apiHealthPath, healthz)

	// route not found
	router.NoRoute(func(ctx *gin.Context) {
		logger.Error(errs.RouteNotFound.String())
		errs.HTTPErrorResponse(ctx, errs.E(errs.Op("NoRoute"), errs.RouteNotFound), &errs.Error{})
	})

	newCommonRoute(ctx, router, provider)
	newStudentRoute(ctx, router, provider)

	// Init swagger routes
	if cfg.Env == config.EnvLocal || cfg.Env == config.EnvDevelopment {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

// healthz for checking service status
func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "welcome to an kms api",
	})
}
