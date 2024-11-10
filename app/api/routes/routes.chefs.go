package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/api/handler/dish"
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
		registry.InjectedClassUseCase(ctx, provider),
		authCookie,
	)
	// auth
	V1AuthRoute := apiV1Group.Group("/auth")
	{
		V1AuthRoute.POST("/refresh", authHdl.Refresh)
		V1AuthRoute.POST("/logout", authHdl.Logout)
		// V1AuthRoute.POST("/change-password", authHdl.ChangePassword)
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
		V1DishRoute.GET("/week", dishHdl.GetDishesForWeek)
	}
}
