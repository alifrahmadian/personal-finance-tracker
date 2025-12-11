package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alifrahmadian/personal-finance-tracker/internal/errors"
)

type AppConfig struct {
	AppPort int
	AppEnv  string
}

func LoadAppConfig() (*AppConfig, error) {
	appPortStr := strings.TrimSpace(os.Getenv("APP_PORT"))
	if appPortStr == "" {
		appPortStr = "8080"
	}
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid APP_PORT %q: %w", appPortStr, errors.ErrFailToParseAppPort)
	}

	appEnv := strings.TrimSpace(os.Getenv("APP_ENV"))
	if appEnv == "" {
		appEnv = APP_ENV_DEVELOPMENT
	}

	return &AppConfig{
		AppPort: appPort,
		AppEnv:  strings.ToLower(appEnv),
	}, nil
}

func ValidateAppConfig(cfg *AppConfig) error {
	if cfg.AppPort <= 0 {
		return errors.ErrAppPortMissing
	}

	if !isAllowedAppEnv(cfg.AppEnv) {
		return errors.ErrInvalidAppEnv
	}

	return nil
}
