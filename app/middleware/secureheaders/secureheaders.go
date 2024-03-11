package secureheaders

import "github.com/gin-gonic/gin"

func SecureHeaders(c *gin.Context) {
	c.Writer.Header().Add("X-Frame-Options", "DENY")
	c.Writer.Header().Add("Cache-Control", "no-store")
	c.Writer.Header().Add("X-Content-Type-Options", "nosniff")
	c.Writer.Header().Add("X-XSS-Protection", "1; mode=block")

	c.Next()
}
