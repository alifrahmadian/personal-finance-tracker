package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Config struct {
	AppConfig     *AppConfig
	DBConfig      *DBConfig
	AuthConfig    *AuthConfig
	LogConfig     *LogConfig
	HandlerConfig *HandlerConfig
}

func LoadConfig() (*Config, *logrus.Logger, func() error, error) {
	appCfg, err := LoadAppConfig()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load app config error: %w", err)
	}

	dbCfg, err := LoadDBConfig()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load db config error: %w", err)
	}

	authCfg, err := LoadAuthConfig()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load auth config error: %w", err)
	}

	logCfg, err := LoadLogConfig()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load log config error: %w", err)
	}

	handlerCfg, err := LoadHandlerConfig()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load handler config error: %w", err)
	}

	if err := validateConfigs(&Config{
		AppConfig:     appCfg,
		DBConfig:      dbCfg,
		AuthConfig:    authCfg,
		LogConfig:     logCfg,
		HandlerConfig: handlerCfg,
	}); err != nil {
		return nil, nil, nil, err
	}

	logger, cleanup, err := SetupLogger(logCfg)
	if err != nil {
		return nil, nil, nil, err
	}

	return &Config{
		AppConfig:     appCfg,
		DBConfig:      dbCfg,
		AuthConfig:    authCfg,
		LogConfig:     logCfg,
		HandlerConfig: handlerCfg,
	}, logger, cleanup, nil
}

func validateConfigs(cfg *Config) error {
	if err := ValidateAppConfig(cfg.AppConfig); err != nil {
		return fmt.Errorf("invalid app config: %w", err)
	}
	if err := ValidateDBConfig(cfg.DBConfig); err != nil {
		return fmt.Errorf("invalid db config: %w", err)
	}
	if err := ValidateAuthConfig(cfg.AuthConfig); err != nil {
		return fmt.Errorf("invalid auth config: %w", err)
	}
	if err := ValidateLogConfig(cfg.LogConfig); err != nil {
		return fmt.Errorf("invalid log config: %w", err)
	}
	if err := ValidateHandlerConfig(cfg.HandlerConfig); err != nil {
		return fmt.Errorf("invalid handler config: %w", err)
	}
	return nil
}
