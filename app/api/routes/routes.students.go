package routes

import (
	"context"
	"kms/app/api/handler/auth"
	"kms/app/api/handler/base"
	"kms/app/api/handler/chat"
	"kms/app/domain/entity"
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

	authCookie := base.NewAuthCookieHandler(provider.Config.Env, entity.AccessKey, entity.RefreshKey, entity.CookiePath, entity.CookiePathRefreshToken, entity.CookieHTTPOnly, entity.CookieMaxAge)

	// auth handler
	authHdl := auth.NewHandler(
		registry.InjectedAuthUseCase(ctx, provider),
		authCookie,
	)
	// auth
	V1AuthRoute := apiV1Group.Group("/auth")
	{
		V1AuthRoute.POST("/refresh", authHdl.Refresh)
		V1AuthRoute.POST("/logout", authHdl.Logout)
		// V1AuthRoute.POST("/change-password", authHdl.ChangePassword)
	}

	// profile
	V1ProfileRoute := apiV1Group.Group("/profile")
	{
		V1ProfileRoute.GET("/me", authHdl.GetProfile)
		V1ProfileRoute.PUT("/me", authHdl.UpdateProfile)
	}

	// chat handler
	chatHdl := chat.NewHandler(
		registry.InjectedChatUseCase(ctx, provider),
		authCookie,
	)
	// chat
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
}
