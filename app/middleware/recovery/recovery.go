package recovery

import (
	"kms/app/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery create a middleware for recovering from panic
func Recovery() gin.HandlerFunc {
	const op errs.Op = "middleware/Recovery"
	// Use nil as writer to prevent Gin to log sensitive information
	// of the request to the default stdout.
	return gin.CustomRecoveryWithWriter(nil, func(ctx *gin.Context, recover interface{}) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{
			"code":    errs.Internal.String(),
			"message": recover,
			"op":      op,
		}})
		ctx.Abort()
	})
}
