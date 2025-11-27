package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alifrahmadian/personal-finance-tracker/internal/errors"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
	DBSSLMode         string
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime time.Duration
}

func LoadDBConfig() (*DBConfig, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, errors.ErrFailToParseDBPort
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbSSLMode == "" {
		dbSSLMode = defaultSSLMode(os.Getenv("APP_ENV"))
	}

	dbMaxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil || dbMaxOpenConns <= 0 {
		dbMaxOpenConns = 5
	}

	dbMaxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil || dbMaxIdleConns <= 0 {
		dbMaxIdleConns = 5
	}

	var dbConnMaxLifetime time.Duration
	lifetimeStr := os.Getenv("DB_CONN_MAX_LIFETIME")
	if lifetimeStr != "" {
		lifetime, err := time.ParseDuration(lifetimeStr)
		if err != nil {
			return nil, fmt.Errorf("invalid DB_CONN_MAX_LIFETIME: %w", err)
		}

		dbConnMaxLifetime = lifetime
	}

	if lifetimeStr == "" {
		dbConnMaxLifetime = 5 * time.Minute
	}

	return &DBConfig{
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            dbPort,
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		DBSSLMode:         dbSSLMode,
		DBMaxOpenConns:    dbMaxOpenConns,
		DBMaxIdleConns:    dbMaxIdleConns,
		DBConnMaxLifetime: dbConnMaxLifetime,
	}, nil
}

func ValidateDBConfig(cfg *DBConfig) error {
	switch {
	case cfg.DBHost == "":
		return errors.ErrMissingDBHost
	case cfg.DBUser == "":
		return errors.ErrMissingDBUser
	case cfg.DBPassword == "":
		return errors.ErrMissingDBPassword
	case cfg.DBName == "":
		return errors.ErrMissingDBName
	case cfg.DBPort <= 0:
		return errors.ErrInvalidDBPort
	}

	if !isAllowedSSLMode(cfg.DBSSLMode) {
		return errors.ErrInvalidDBSSLMode
	}

	return nil
}

func ConnectDB(cfg *DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	applyPoolSettings(db, cfg)
	return db, nil
}
