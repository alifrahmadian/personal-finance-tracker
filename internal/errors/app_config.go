package errors

import "errors"

var (
	ErrAppPortMissing     = errors.New("application port is missing")
	ErrFailToParseAppPort = errors.New("failed to parse application port")
	ErrAppEnvMissing      = errors.New("application environment is missing")
	ErrInvalidAppEnv      = errors.New("invalid application environment")
)
