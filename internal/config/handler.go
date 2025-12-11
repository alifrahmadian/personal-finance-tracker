package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alifrahmadian/personal-finance-tracker/internal/errors"
)

type HandlerConfig struct {
	BasePath          string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	EnableHealthCheck bool
	EnableSwagger     bool
	RequestLogger     bool
	AllowedOrigins    []string
}

func LoadHandlerConfig() (*HandlerConfig, error) {
	basePath := strings.TrimSpace(os.Getenv("BASE_PATH"))
	if basePath == "" {
		return nil, errors.ErrBasePathMissing
	}

	readTimeoutStr := strings.TrimSpace(os.Getenv("READ_TIMEOUT"))
	if readTimeoutStr == "" {
		return nil, errors.ErrReadTimeoutMissing
	}

	readTimeout, err := time.ParseDuration(readTimeoutStr)
	if err != nil {
		return nil, err
	}

	writeTimeoutStr := strings.TrimSpace(os.Getenv("WRITE_TIMEOUT"))
	if writeTimeoutStr == "" {
		return nil, errors.ErrWriteTimeoutMissing
	}

	writeTimeout, err := time.ParseDuration(writeTimeoutStr)
	if err != nil {
		return nil, err
	}

	idleTimeoutStr := strings.TrimSpace(os.Getenv("IDLE_TIMEOUT"))
	if idleTimeoutStr == "" {
		return nil, errors.ErrIdleTimeoutMissing
	}

	idleTimeout, err := time.ParseDuration(idleTimeoutStr)
	if err != nil {
		return nil, err
	}

	enableHealthCheckStr := strings.TrimSpace(os.Getenv("ENABLE_HEALTH_CHECK"))
	if enableHealthCheckStr == "" {
		return nil, errors.ErrEnableHealthCheckMissing
	}

	enableHealthCheck, err := strconv.ParseBool(enableHealthCheckStr)
	if err != nil {
		return nil, err
	}

	enableSwaggerStr := strings.TrimSpace(os.Getenv("ENABLE_SWAGGER"))
	if enableSwaggerStr == "" {
		return nil, errors.ErrEnableSwaggerMissing
	}

	enableSwagger, err := strconv.ParseBool(enableSwaggerStr)
	if err != nil {
		return nil, err
	}

	requestLoggerStr := strings.TrimSpace(os.Getenv("REQUEST_LOGGER"))
	if requestLoggerStr == "" {
		return nil, errors.ErrRequestLoggerMissing
	}

	requestLogger, err := strconv.ParseBool(requestLoggerStr)
	if err != nil {
		return nil, err
	}

	allowedOriginsStr := strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS"))
	if allowedOriginsStr == "" {
		return nil, errors.ErrAllowedOriginsMissing
	}

	rawOrigins := strings.Split(strings.TrimSpace(allowedOriginsStr), ",")

	allowedOrigins := make([]string, 0, len(rawOrigins))
	for _, origin := range rawOrigins {
		trimmedOrigin := strings.TrimSpace(origin)
		if trimmedOrigin != "" {
			allowedOrigins = append(allowedOrigins, trimmedOrigin)
		}
	}

	return &HandlerConfig{
		BasePath:          basePath,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		EnableHealthCheck: enableHealthCheck,
		EnableSwagger:     enableSwagger,
		RequestLogger:     requestLogger,
		AllowedOrigins:    allowedOrigins,
	}, nil
}

func ValidateHandlerConfig(cfg *HandlerConfig) error {
	// Add validation logic if needed
	if !strings.HasPrefix(cfg.BasePath, "/") {
		return errors.ErrInvalidBasePathNoSlash
	}

	if cfg.ReadTimeout <= 0 {
		return errors.ErrReadTimeoutZero
	}

	if cfg.WriteTimeout <= 0 {
		return errors.ErrWriteTimeoutZero
	}

	if cfg.IdleTimeout <= 0 {
		return errors.ErrIdleTimeoutZero
	}

	if cfg.ReadTimeout > HANDLER_MAX_READ_TIMEOUT*time.Second {
		return errors.ErrMaxReadTimeoutExceeded
	}

	if cfg.WriteTimeout > HANDLER_MAX_WRITE_TIMEOUT*time.Second {
		return errors.ErrMaxWriteTimeoutExceeded
	}

	if cfg.IdleTimeout > HANDLER_MAX_IDLE_TIMEOUT*time.Second {
		return errors.ErrMaxIdleTimeoutExceeded
	}

	if len(cfg.AllowedOrigins) == 0 {
		return errors.ErrAllowedOriginsMissing
	}

	if !containsWildcardOrigin(cfg.AllowedOrigins) {
		for _, origin := range cfg.AllowedOrigins {
			if !isValidOrigin(origin) {
				return errors.ErrHandlerInvalidOrigin
			}
		}
	}

	return nil
}
