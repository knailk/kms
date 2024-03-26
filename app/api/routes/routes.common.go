package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

const (
	accessKey              = "KMS_jwt_access"
	refreshKey             = "KMS_jwt_refresh"
	cookiePath             = "/api/v1/customer"
	cookiePathRefreshToken = "/api/v1/auth/refresh"
	cookieHTTPOnly         = false
	cookieMaxAge           = 14400
)

func newCommonRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
) {
	apiV1Group := router.Group(apiCommonV1)

	authCookie := base.NewAuthCookieHandler(provider.Config.Env, accessKey, refreshKey, cookiePath, cookiePathRefreshToken, cookieHTTPOnly, cookieMaxAge)

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
