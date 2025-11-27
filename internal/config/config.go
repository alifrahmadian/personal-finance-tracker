package config

import "database/sql"

type Config struct {
	DB            *sql.DB
	AppConfig     *AppConfig
	DBConfig      *DBConfig
	AuthConfig    *AuthConfig
	LogConfig     *LogConfig
	HandlerConfig *HandlerConfig
}
