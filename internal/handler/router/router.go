package router

import (
	"github.com/alifrahmadian/personal-finance-tracker/internal/config"
	"github.com/alifrahmadian/personal-finance-tracker/internal/handler/health"
	"github.com/alifrahmadian/personal-finance-tracker/internal/handler/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Health *health.Handler
}

func NewRouter(handlerCfg *config.HandlerConfig, logger *logrus.Logger, handlers *Handlers) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())

	if handlerCfg.RequestLogger {
		router.Use(middleware.RequestLogger(logger))
	}

	api := router.Group(handlerCfg.BasePath)

	if handlerCfg.EnableHealthCheck {
		health.RegisterRoutes(api, handlers.Health)
	}

	return router
}
