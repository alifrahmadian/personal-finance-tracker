package config

import (
	"os"
	"strings"

	"github.com/alifrahmadian/personal-finance-tracker/internal/errors"
	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	LogLevel    string
	LogFormat   string
	LogOutput   string
	LogFilePath string
}

func LoadLogConfig() (*LogConfig, error) {
	logLevel := strings.TrimSpace(strings.ToLower(os.Getenv("LOG_LEVEL")))
	logFormat := strings.TrimSpace(strings.ToLower(os.Getenv("LOG_FORMAT")))
	logOutput := strings.TrimSpace(strings.ToLower(os.Getenv("LOG_OUTPUT")))
	logFilePath := strings.TrimSpace(os.Getenv("LOG_FILE_PATH"))

	if logLevel == "" {
		return nil, errors.ErrMissingLogLevel
	}
	if logFormat == "" {
		return nil, errors.ErrMissingLogFormat
	}
	if logOutput == "" {
		return nil, errors.ErrMissingLogOutput
	}

	return &LogConfig{
		LogLevel:    logLevel,
		LogFormat:   logFormat,
		LogOutput:   logOutput,
		LogFilePath: logFilePath,
	}, nil
}

func ValidateLogConfig(cfg *LogConfig) error {
	// Add validation logic if needed
	if !isAllowedLogLevel(cfg.LogLevel) {
		return errors.ErrInvalidLogLevel
	}
	if !isAllowedLogFormat(cfg.LogFormat) {
		return errors.ErrInvalidLogFormat
	}
	if !isAllowedLogOutput(cfg.LogOutput) {
		return errors.ErrInvalidLogOutput
	}

	if cfg.LogOutput == LOG_OUTPUT_FILE && cfg.LogFilePath == "" {
		return errors.ErrLogFilePathMissing
	}

	return nil
}

func SetupLogger(cfg *LogConfig) (*logrus.Logger, func() error, error) {
	// Implement logger setup based on cfg
	if cfg == nil {
		return nil, nil, errors.ErrLogConfigNotFound
	}

	validLogConfig := ValidateLogConfig(cfg)
	if validLogConfig != nil {
		return nil, nil, validLogConfig
	}

	logger := logrus.New()

	parsedLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, nil, err
	}

	logger.SetLevel(parsedLevel)
	logger.SetFormatter(resolveLogFormatter(cfg.LogFormat))

	writer, cleanup, err := resolveLogOutput(cfg)
	if err != nil {
		return nil, nil, err
	}
	logger.SetOutput(writer)

	return logger, cleanup, nil
}
