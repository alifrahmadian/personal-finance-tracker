package errors

import "errors"

var (
	ErrDBConnectionFailed = errors.New("database connection failed")
	ErrFailToParseDBPort  = errors.New("failed to parse database port")
	ErrInvalidDBPort      = errors.New("invalid database port")
	ErrMissingDBHost      = errors.New("missing database host")
	ErrMissingDBPort      = errors.New("missing database port")
	ErrMissingDBUser      = errors.New("missing database username")
	ErrMissingDBName      = errors.New("missing database name")
	ErrMissingDBPassword  = errors.New("missing database password")
	ErrInvalidDBSSLMode   = errors.New("invalid database SSL mode")
)
