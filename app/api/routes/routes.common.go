package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/registry"
	"kms/app/domain/entity"

	"github.com/gin-gonic/gin"
)

func newCommonRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
) {
	apiV1Group := router.Group(apiCommonV1)

	authCookie := base.NewAuthCookieHandler(provider.Config.Env, entity.AccessKey, entity.RefreshKey, entity.CookiePath, entity.CookiePathRefreshToken, entity.CookieHTTPOnly, entity.CookieMaxAge)

	// auth
	authHdl := auth.NewHandler(
		registry.InjectedAuthUseCase(ctx, provider),
		authCookie,
	)

	V1AuthRoute := apiV1Group.Group("/auth")
	V1AuthRoute.POST("/login", authHdl.Login)
	// V1AuthRoute.POST("/forgot-password/request", sessionHdl.ForgotPassword)
	// V1AuthRoute.POST("/forgot-password/confirm", sessionHdl.ConfirmForgotPassword)
}
