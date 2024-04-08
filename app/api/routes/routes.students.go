package routes

import (
	"context"
	"kms/app/api/handler/base"
	"kms/app/api/handler/user"
	"kms/app/domain/entity"
	"kms/app/middleware/author"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

func newStudentRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
	authCookie base.AuthCookieHandler,
) {
	apiV1Group := router.Group(apiStudentV1)
	apiV1Group.Use(author.NewAuthMiddleware(entity.UserTypeStudent))

	// user handler
	userHdl := user.NewHandler(
		registry.InjectedUserUseCase(ctx, provider),
		authCookie,
	)
	V1UserRoute := apiV1Group.Group("/profile")
	{
		V1UserRoute.GET("/me", userHdl.GetProfile)
		V1UserRoute.PUT("/me", userHdl.UpdateUser)
		V1UserRoute.GET("/", userHdl.SearchUser)
	}
}
