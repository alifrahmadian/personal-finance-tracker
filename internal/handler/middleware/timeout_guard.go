package middleware

import (
	"context"
	"net/http"

	"github.com/alifrahmadian/personal-finance-tracker/internal/config"
	"github.com/gin-gonic/gin"
)

func TimeoutGuard(handlerConfig *config.HandlerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), handlerConfig.WriteTimeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()

		if ctx.Err() == context.DeadlineExceeded && !c.IsAborted() {
			c.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}
	}
}
