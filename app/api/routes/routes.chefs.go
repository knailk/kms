package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/domain/entity"
	"kms/app/middleware/author"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

func newChefRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
	authCookie base.AuthCookieHandler,
) {
	apiV1Group := router.Group(apiChefV1)
	apiV1Group.Use(author.NewAuthMiddleware(entity.UserTypeChef))

	// auth handler
	authHdl := auth.NewHandler(
		registry.InjectedAuthUseCase(ctx, provider),
		registry.InjectedClassUseCase(ctx,provider),
		authCookie,
	)
	// auth
	V1AuthRoute := apiV1Group.Group("/auth")
	{
		V1AuthRoute.POST("/refresh", authHdl.Refresh)
		V1AuthRoute.POST("/logout", authHdl.Logout)
		// V1AuthRoute.POST("/change-password", authHdl.ChangePassword)
	}
}
