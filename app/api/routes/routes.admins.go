package routes

import (
	"context"
	"kms/app/api/handler/base"
	"kms/app/api/handler/class"
	"kms/app/domain/entity"
	"kms/app/middleware/author"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

func newAdminRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
	authCookie base.AuthCookieHandler,
) {
	apiV1Group := router.Group(apiAdminV1)
	apiV1Group.Use(author.NewAuthMiddleware(entity.UserTypeAdmin))

	classHdl := class.NewHandler(
		registry.InjectedClassUseCase(ctx, provider),
		authCookie,
	)
	V1ClassRoute := apiV1Group.Group("/class")
	{
		V1ClassRoute.GET("", classHdl.ListClasses)
		V1ClassRoute.POST("", classHdl.CreateClass)
		V1ClassRoute.PUT("/:id", classHdl.UpdateClass)
		V1ClassRoute.DELETE("/:id", classHdl.DeleteClass)
		V1ClassRoute.POST("/:id/members", classHdl.AddMembersToClass)
		V1ClassRoute.DELETE("/:id/members", classHdl.RemoveMembersFromClass)
	}
}
