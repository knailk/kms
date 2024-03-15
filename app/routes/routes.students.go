package routes

import (
	"context"
	"kms/app/entity"
	"kms/app/middleware/author"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

func newStudentRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
) {
	apiV1Group := router.Group(apiStudentV1)
	apiV1Group.Use(author.NewAuthMiddleware(entity.UserTypeStudent))

	// auth
	V1AuthRoute := apiV1Group.Group("/auth")
	V1AuthRoute.GET("/me", sessionHdl.GetUserInfo)
	V1AuthRoute.POST("/change-password", sessionHdl.ChangePassword)
	V1AuthRoute.POST("/refresh", sessionHdl.Refresh)
	V1AuthRoute.POST("/logout", sessionHdl.Logout)
}
