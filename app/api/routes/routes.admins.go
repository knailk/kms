package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/api/handler/class"
	"kms/app/api/handler/dish"
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

	authHdl := auth.NewHandler(
		registry.InjectedAuthUseCase(ctx, provider),
		registry.InjectedClassUseCase(ctx, provider),
		authCookie,
	)
	{
		V1AuthTokenRoute := apiV1Group.Group("/auth")
		V1AuthTokenRoute.POST("/register-confirm", authHdl.RegisterConfirm)
		V1AuthTokenRoute.GET("/register-request-list", authHdl.RegisterRequestList)
	}

	classHdl := class.NewHandler(
		registry.InjectedClassUseCase(ctx, provider),
		authCookie,
	)
	{
		V1ClassRoute := apiV1Group.Group("/class")
		V1ClassRoute.GET("", classHdl.ListClasses)
		V1ClassRoute.POST("", classHdl.CreateClass)
		V1ClassRoute.PUT("/:id", classHdl.UpdateClass)
		V1ClassRoute.DELETE("/:id", classHdl.DeleteClass)
		V1ClassRoute.POST("/:id/members", classHdl.AddMembersToClass)
		V1ClassRoute.DELETE("/:id/members", classHdl.RemoveMembersFromClass)

		V1ClassRoute.POST("/schedules", classHdl.CreateSchedule)
		V1ClassRoute.PUT("/schedules/:id", classHdl.UpdateSchedule)
		V1ClassRoute.DELETE("/schedules/:id", classHdl.DeleteSchedule)
	}

	dishHdl := dish.NewHandler(
		registry.InjectedDishUseCase(ctx, provider),
		authCookie,
	)
	//dish
	{
		V1DishRoute := apiV1Group.Group("/dishes")
		V1DishRoute.POST("", dishHdl.CreateDish)
		V1DishRoute.PUT("/:id", dishHdl.UpdateDish)
		V1DishRoute.DELETE("/:id", dishHdl.DeleteDish)
	}
}
