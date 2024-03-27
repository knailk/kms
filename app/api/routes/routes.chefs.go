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
) {
	apiV1Group := router.Group(apiChefV1)
	apiV1Group.Use(author.NewAuthMiddleware(entity.UserTypeChef))

	authCookie := base.NewAuthCookieHandler(provider.Config.Env, entity.AccessKey, entity.RefreshKey, entity.CookiePath, entity.CookiePathRefreshToken, entity.CookieHTTPOnly, entity.CookieMaxAge)

	// auth handler
	authHdl := auth.NewHandler(
		registry.InjectedAuthUseCase(ctx, provider),
		authCookie,
	)
	// auth
	V1AuthRoute := apiV1Group.Group("/auth")
	{
		V1AuthRoute.GET("/me", authHdl.GetInfo)
		V1AuthRoute.POST("/refresh", authHdl.Refresh)
		V1AuthRoute.POST("/logout", authHdl.Logout)
		// V1AuthRoute.POST("/change-password", authHdl.ChangePassword)
	}
}
