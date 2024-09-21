package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Function to log HTTP requests using logrus
func HTTPLogger(param gin.LogFormatterParams) string {
	log := logrus.WithFields(logrus.Fields{
		"time":       param.TimeStamp.Format(time.RFC3339),
		"status":     param.StatusCode,
		"method":     param.Method,
		"path":       param.Path,
		"latency":    param.Latency,
		"client_ip":  param.ClientIP,
		"error":      param.ErrorMessage,
		"user_agent": param.Request.UserAgent(),
	})

	// Determine log level based on status code
	switch {
	case param.StatusCode >= 500:
		log.Error("[LOGGING HTTP] Internal Server Error")
	case param.StatusCode >= 400:
		log.Warn("[LOGGING HTTP] Client Error")
	default:
		log.Info("[LOGGING HTTP] Success")
	}

	return ""
}
