package config

import "database/sql"

func isAllowedSSLMode(mode string) bool {
	allowedModes := map[string]bool{
		"disable":     true,
		"require":     true,
		"verify-ca":   true,
		"verify-full": true,
	}
	return allowedModes[mode]
}

func defaultSSLMode(appEnv string) string {
	if appEnv == "production" {
		return "verify-full"
	}
	return "disable"
}

func applyPoolSettings(db *sql.DB, cfg *DBConfig) {
	if cfg.DBMaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	}
	if cfg.DBMaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	}
	if cfg.DBConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.DBConnMaxLifetime)
	}
}
