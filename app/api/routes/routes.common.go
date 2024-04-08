package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/api/handler/chat"
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
		authCookie,
	)

	V1AuthRoute := apiV1Group.Group("/auth")
	V1AuthRoute.POST("/login", authHdl.Login)

	apiV1Group.Use(author.NewAuthMiddleware(entity.UserTypeCommon))

	V1AuthTokenRoute := apiV1Group.Group("/auth")
	{
		V1AuthTokenRoute.POST("/refresh", authHdl.Refresh)
		V1AuthTokenRoute.POST("/logout", authHdl.Logout)
		// V1AuthRoute.POST("/change-password", authHdl.ChangePassword)
	}

	// chat handler
	chatHdl := chat.NewHandler(
		registry.InjectedChatUseCase(ctx, provider),
		authCookie,
	)
	V1ChatRoute := apiV1Group.Group("/chat")
	{
		V1ChatRoute.POST("/", chatHdl.CreateChat)
		V1ChatRoute.PUT("/member", chatHdl.AddMember)
		V1ChatRoute.GET("/", chatHdl.ListChats)
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
	V1UserRoute := apiV1Group.Group("/profile")
	{
		V1UserRoute.GET("/me", userHdl.GetProfile)
		V1UserRoute.PUT("/me", userHdl.UpdateUser)
		V1UserRoute.GET("/", userHdl.SearchUser)
	}

	// classHdl := class

}
