package health

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	healthGroup := rg.Group("/health")
	{
		healthGroup.GET("/", h.CheckHealth)
	}
}
