package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/api/handler/chat"
	"kms/app/api/handler/class"
	"kms/app/api/handler/user"
	"kms/app/domain/entity"
	"kms/app/middleware/author"
	"kms/app/registry"

	"github.com/gin-gonic/gin"
)

func newCommonRoute(
	ctx context.Context,
	router *gin.Engine,
	provider *registry.Provider,
	authCookie base.AuthCookieHandler,
) {
	apiV1Group := router.Group(apiCommonV1)
	// auth
	authHdl := auth.NewHandler(
		registry.InjectedAuthUseCase(ctx, provider),
		registry.InjectedClassUseCase(ctx, provider),
		authCookie)
	{
		V1AuthRoute := apiV1Group.Group("/auth")
		V1AuthRoute.POST("/login", authHdl.Login)
		V1AuthRoute.POST("/register", authHdl.RegisterRequest)
	}

	classHdl := class.NewHandler(
		registry.InjectedClassUseCase(ctx, provider),
		authCookie,
	)
	{
		V1ClassRoute := apiV1Group.Group("/classes")
		V1ClassRoute.GET("", classHdl.ListClasses)
	}

	// ------------------------------------ BEFORE LOGIN ------------------------------------

	apiV1Group.Use(author.NewAuthMiddleware(entity.UserTypeCommon))

	// ------------------------------------ AFTER LOGIN ------------------------------------

	{
		V1AuthTokenRoute := apiV1Group.Group("/auth")
		V1AuthTokenRoute.POST("/refresh", authHdl.Refresh)
		V1AuthTokenRoute.POST("/logout", authHdl.Logout)
		// V1AuthRoute.POST("/change-password", authHdl.ChangePassword)
	}
	{
		chatHdl := chat.NewHandler(
			registry.InjectedChatUseCase(ctx, provider),
			authCookie,
		)
		V1ChatRoute := apiV1Group.Group("/chat")
		V1ChatRoute.POST("", chatHdl.CreateChat)
		V1ChatRoute.PUT("/member", chatHdl.AddMember)
		V1ChatRoute.GET("", chatHdl.ListChats)
		V1ChatRoute.GET("/:id", chatHdl.GetChat)
		V1ChatRoute.PUT("/:id", chatHdl.UpdateChat)
		V1ChatRoute.DELETE("/:id", chatHdl.DeleteChat)
		V1ChatRoute.POST("/:id/message", chatHdl.CreateMessage)
	}
	// user handler
	userHdl := user.NewHandler(
		registry.InjectedUserUseCase(ctx, provider),
		authCookie,
	)
	{
		V1UserRoute := apiV1Group.Group("/profile")
		V1UserRoute.GET("/me", userHdl.GetProfile)
		V1UserRoute.PUT("/me", userHdl.UpdateUser)
		V1UserRoute.GET("", userHdl.SearchUser)
		V1UserRoute.GET("/teacher-available", userHdl.ListTeachersAvailable)
		V1UserRoute.GET("/driver-available", userHdl.ListDriversAvailable)
	}
	{
		V1ClassRoute := apiV1Group.Group("/class")
		V1ClassRoute.GET("/me", classHdl.GetClass)
		V1ClassRoute.GET("/:id/members", classHdl.ListMembersInClass)
		V1ClassRoute.POST("/:id/members", classHdl.CheckInOut)
	}
}
