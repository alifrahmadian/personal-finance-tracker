package main

import (
	"fmt"

	"github.com/alifrahmadian/personal-finance-tracker/internal/config"
	"github.com/alifrahmadian/personal-finance-tracker/internal/handler/health"
	"github.com/alifrahmadian/personal-finance-tracker/internal/handler/router"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *config.Config
}

func newApp() *App {
	cfg, logger, cleanup, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if cfg.AppConfig.AppEnv == config.APP_ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	healthHandler := health.NewHandler(logger)

	router := router.NewRouter(cfg.HandlerConfig, logger, &router.Handlers{
		Health: healthHandler,
	})

	return &App{
		Router: router,
		Config: cfg,
	}
}

func main() {
	app := newApp()
	addr := fmt.Sprintf(":%d", app.Config.AppConfig.AppPort)

	err := app.Router.Run(addr)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
