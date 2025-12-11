package errors

import "errors"

var (
	ErrInvalidAppConfig     = errors.New("invalid app config")
	ErrInvalidDBConfig      = errors.New("invalid db config")
	ErrInvalidAuthConfig    = errors.New("invalid auth config")
	ErrInvalidLogConfig     = errors.New("invalid log config")
	ErrInvalidHandlerConfig = errors.New("invalid handler config")
)
