package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alifrahmadian/personal-finance-tracker/internal/errors"
)

type AuthConfig struct {
	JWTSecret    string
	JWTExpiresIn time.Duration
}

func LoadAuthConfig() (*AuthConfig, error) {
	jwtSecret := strings.TrimSpace(os.Getenv("JWT_SECRET"))

	expiryStr := strings.TrimSpace(os.Getenv("JWT_EXPIRES_IN"))
	if expiryStr == "" {
		expiryStr = "3600" // default  to 1 hour when not set
	}

	expiry, err := strconv.Atoi(expiryStr)
	if err != nil {
		return nil, errors.ErrFailToParseJWTExpiry
	}

	if expiry <= 0 {
		expiry = 3600
	}

	expiryDuration := time.Duration(expiry) * time.Second

	return &AuthConfig{
		JWTSecret:    jwtSecret,
		JWTExpiresIn: expiryDuration,
	}, nil
}

func ValidateAuthConfig(cfg *AuthConfig) error {
	if cfg.JWTSecret == "" {
		return errors.ErrJWTSecretMissing
	}

	if cfg.JWTExpiresIn <= 0 {
		return errors.ErrInvalidJWTExpiry
	}

	return nil
}
