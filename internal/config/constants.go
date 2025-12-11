package config

const (
	APP_ENV_DEVELOPMENT = "development"
	APP_ENV_PRODUCTION  = "production"
	APP_ENV_STAGING     = "staging"
)

const (
	LOG_LEVEL_DEBUG = "debug"
	LOG_LEVEL_INFO  = "info"
	LOG_LEVEL_WARN  = "warn"
	LOG_LEVEL_ERROR = "error"
)

const (
	LOG_FORMAT_JSON = "json"
	LOG_FORMAT_TEXT = "text"
)

const (
	LOG_OUTPUT_STDOUT = "stdout"
	LOG_OUTPUT_STDERR = "stderr"
	LOG_OUTPUT_FILE   = "file"
)

const (
	HANDLER_MAX_READ_TIMEOUT  = 120
	HANDLER_MAX_WRITE_TIMEOUT = 120
	HANDLER_MAX_IDLE_TIMEOUT  = 300
)
