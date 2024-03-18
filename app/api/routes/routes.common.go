package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

func newCommonRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
) {
	apiV1Group := router.Group(apiCommonV1)

	// auth
	authHdl := auth.NewHandler(
		registry.InjectedAuthUseCase(ctx, provider),
	)

	V1AuthRoute := apiV1Group.Group("/auth")
	V1AuthRoute.POST("/sign-in", authHdl.Login)
	// V1AuthRoute.POST("/forgot-password/request", sessionHdl.ForgotPassword)
	// V1AuthRoute.POST("/forgot-password/confirm", sessionHdl.ConfirmForgotPassword)

}
