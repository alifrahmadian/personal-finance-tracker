package errors

import "errors"

var (
	ErrInvalidLogLevel    = errors.New("invalid log level")
	ErrInvalidLogFormat   = errors.New("invalid log format")
	ErrInvalidLogOutput   = errors.New("invalid log output")
	ErrLogConfigNotFound  = errors.New("log configuration not found")
	ErrLogFilePathMissing = errors.New("log file path missing")
	ErrMissingLogLevel    = errors.New("log level is missing")
	ErrMissingLogFormat   = errors.New("log format is missing")
	ErrMissingLogOutput   = errors.New("log output is missing")
)
