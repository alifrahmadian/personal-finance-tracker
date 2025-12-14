package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/alifrahmadian/personal-finance-tracker/internal/config"
	"github.com/gin-gonic/gin"
)

func CORS(cfg *config.HandlerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		if len(cfg.AllowedOrigins) == 0 {
			c.Next()
			return
		}

		if allowedContainsStar(cfg.AllowedOrigins) {
			c.Header("Access-Control-Allow-Origin", "*")
		} else if originAllowed(origin, cfg.AllowedOrigins) {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "CORS origin not allowed",
			})
			return
		}

		c.Header("Access-Control-Allow-Origin", cfg.AllowedOrigins[0])
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,X-Request-ID")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func allowedContainsStar(allowedOrigins []string) bool {
	return slices.Contains(allowedOrigins, "*")
}

func originAllowed(origin string, allowedOrigins []string) bool {
	if origin == "" {
		return false
	}

	for _, ao := range allowedOrigins {
		if origin == strings.TrimSpace(ao) {
			return true
		}
	}

	return false
}
