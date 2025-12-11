package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger *logrus.Logger
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) CheckHealth(c *gin.Context) {
	h.logger.Info("Health check endpoint called")
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
