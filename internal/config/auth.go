package config

import (
	"errors"
	"os"
	"strconv"
)

type AuthConfig struct {
	JWTSecret string
	JWTExpiry int
}

func LoadAuthConfig() (*AuthConfig, error) {
	expiry, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		return nil, errors.New("invalid JWT_EXPIRES_IN: " + err.Error())
	}

	return &AuthConfig{
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpiry: expiry,
	}, nil
}
