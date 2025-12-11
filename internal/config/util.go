package config

import (
	"database/sql"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func isAllowedSSLMode(mode string) bool {
	allowedModes := map[string]bool{
		"disable":     true,
		"require":     true,
		"verify-ca":   true,
		"verify-full": true,
	}
	return allowedModes[mode]
}

func isAllowedAppEnv(env string) bool {
	allowedEnvs := map[string]bool{
		"development": true,
		"staging":     true,
		"production":  true,
	}
	return allowedEnvs[env]
}

func isAllowedLogLevel(level string) bool {
	allowedLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	return allowedLevels[level]
}

func isAllowedLogFormat(format string) bool {
	allowedFormats := map[string]bool{
		"json": true,
		"text": true,
	}
	return allowedFormats[format]
}

func isAllowedLogOutput(output string) bool {
	allowedOutputs := map[string]bool{
		"stdout": true,
		"stderr": true,
		"file":   true,
	}
	return allowedOutputs[output]
}

func defaultSSLMode(appEnv string) string {
	if appEnv == "production" {
		return "verify-full"
	}
	return "disable"
}

func applyPoolSettings(db *sql.DB, cfg *DBConfig) {
	if cfg.DBMaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	}
	if cfg.DBMaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	}
	if cfg.DBConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.DBConnMaxLifetime)
	}
}

func resolveLogFormatter(format string) logrus.Formatter {
	if format == LOG_FORMAT_JSON {
		return &logrus.JSONFormatter{}
	}

	return &logrus.TextFormatter{FullTimestamp: true}
}

func resolveLogOutput(cfg *LogConfig) (io.Writer, func() error, error) {
	noop := func() error { return nil }

	switch cfg.LogOutput {
	case LOG_OUTPUT_STDOUT:
		return os.Stdout, noop, nil
	case LOG_OUTPUT_STDERR:
		return os.Stderr, noop, nil
	case LOG_OUTPUT_FILE:
		path := strings.TrimSpace(cfg.LogFilePath)
		dir := filepath.Dir(path)
		if dir != "" && dir != "." {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return nil, nil, err
			}
		}
		file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			return nil, nil, err
		}
		cleanup := func() error { return file.Close() }
		return file, cleanup, nil
	default:
		return os.Stdout, noop, nil
	}
}

func containsWildcardOrigin(origins []string) bool {
	for _, origin := range origins {
		if origin == "*" {
			return true
		}
	}
	return false
}

func isValidOrigin(origin string) bool {
	u, err := url.ParseRequestURI(origin)
	if err != nil {
		return false
	}

	if u.Scheme == "" || u.Host == "" {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
}
