package middleware

import (
	"github.com/alifrahmadian/personal-finance-tracker/internal/config"
	"github.com/gin-gonic/gin"
)

func SecurityHeaders(appCfg *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "no-referrer")
		c.Header("Permissions-Policy", "geolocation=()")
		c.Header("Content-Security-Policy", "default-src 'self'")

		if appCfg.AppEnv == config.APP_ENV_PRODUCTION {
			c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		}

		c.Next()
	}
}
