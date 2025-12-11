package errors

import "errors"

var (
	ErrBasePathMissing          = errors.New("base path is missing")
	ErrReadTimeoutMissing       = errors.New("read timeout is missing")
	ErrWriteTimeoutMissing      = errors.New("write timeout is missing")
	ErrIdleTimeoutMissing       = errors.New("idle timeout is missing")
	ErrEnableHealthCheckMissing = errors.New("enable health check flag is missing")
	ErrEnableSwaggerMissing     = errors.New("enable swagger flag is missing")
	ErrRequestLoggerMissing     = errors.New("request logger configuration is missing")
	ErrAllowedOriginsMissing    = errors.New("allowed origins are missing")
	ErrInvalidBasePathNoSlash   = errors.New("base path must start with a slash")
	ErrMaxReadTimeoutExceeded   = errors.New("read timeout exceeds maximum allowed value")
	ErrMaxWriteTimeoutExceeded  = errors.New("write timeout exceeds maximum allowed value")
	ErrMaxIdleTimeoutExceeded   = errors.New("idle timeout exceeds maximum allowed value")
	ErrHandlerInvalidOrigin     = errors.New("invalid origin in allowed origins")
	ErrReadTimeoutZero          = errors.New("read timeout must be greater than zero")
	ErrWriteTimeoutZero         = errors.New("write timeout must be greater than zero")
	ErrIdleTimeoutZero          = errors.New("idle timeout must be greater than zero")
)
