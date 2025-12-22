package errors

import "errors"

var (
	ErrNoAuthHeader            = errors.New("authorization header required")
	ErrInvalidTokenFormat      = errors.New("invalid token format")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidToken            = errors.New("invalid token")
	ErrTokenExpires            = errors.New("token has expired")
)
