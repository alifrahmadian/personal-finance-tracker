package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		requestID := c.GetString(RequestIDKey)
		clientIP := c.ClientIP()

		entry := logger.WithFields(logrus.Fields{
			"method":     method,
			"path":       path,
			"status":     statusCode,
			"latency_ms": latency.Milliseconds(),
			"client_ip":  clientIP,
			"request_id": requestID,
		})

		switch {
		case statusCode >= 500:
			entry.Error("Server error occurred")
		case statusCode >= 400:
			entry.Warn("Client error occurred")
		default:
			entry.Info("Request processed successfully")
		}

	}
}
