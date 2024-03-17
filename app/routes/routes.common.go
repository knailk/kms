package routes

import (
	"context"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

func newCommonRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
) {
	// apiV1Group := router.Group(apiCommonV1)

	// // auth
	// V1AuthRoute := apiV1Group.Group("/auth")
	// V1AuthRoute.POST("/sign-in", sessionHdl.SignIn)
	// V1AuthRoute.POST("/forgot-password/request", sessionHdl.ForgotPassword)
	// V1AuthRoute.POST("/forgot-password/confirm", sessionHdl.ConfirmForgotPassword)

}
