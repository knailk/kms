package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	corsCfg := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Length", "Origin", "cookie", "access-control-allow-origin",
			"authorization, origin, content-type, accept", "X-CSRF-Token", "Pragma"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(corsCfg)
}
