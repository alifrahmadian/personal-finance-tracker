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

func NewRouter(
	handlerCfg *config.HandlerConfig,
	appCfg *config.AppConfig,
	logger *logrus.Logger,
	handlers *Handlers,
) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.SecurityHeaders(appCfg))
	router.Use(middleware.CORS(handlerCfg))
	router.Use(middleware.TimeoutGuard(handlerCfg))

	if handlerCfg.RequestLogger {
		router.Use(middleware.RequestLogger(logger))
	}

	api := router.Group(handlerCfg.BasePath)

	if handlerCfg.EnableHealthCheck {
		health.RegisterRoutes(api, handlers.Health)
	}

	return router
}
