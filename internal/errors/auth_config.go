package errors

import "errors"

var (
	ErrJWTSecretMissing     = errors.New("jwt secret key is missing")
	ErrFailToParseJWTExpiry = errors.New("failed to parse jwt expiry duration")
	ErrInvalidJWTExpiry     = errors.New("invalid jwt expiry duration")
)
