package router

import (
	"github.com/alifrahmadian/personal-finance-tracker/internal/config"
	"github.com/alifrahmadian/personal-finance-tracker/internal/handler/health"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Health *health.Handler
}

func NewRouter(handlerCfg *config.HandlerConfig, logger *logrus.Logger, handlers *Handlers) *gin.Engine {
	router := gin.New()
	api := router.Group(handlerCfg.BasePath)

	health.RegisterRoutes(api, handlers.Health)

	return router
}
