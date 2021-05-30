package config

func GlobalEnvUpdateMap() map[string]string {
	m := make(map[string]string)

	m["AppName"] = "APP_NAME"
	m["AppVersion"] = "APP_VERSION"
	m["ServerHost"] = "SERVER_HOST"
	m["ServerPort"] = "SERVER_PORT"
	m["Environment"] = "APP_ENVIRONMENT"
	m["LogConfFile"] = "LOG_CONF_FILE"
	m["SuperKey"] = "SUPER_KEY"
	m["LogConfig.Level"] = "LOG_CONF_LEVEL"
	m["LogConfig.DevelopmentEnv"] = "LOG_CONF_DEVELOPMENT_ENV"
	m["LogConfig.Encoding"] = "LOG_CONF_ENCODING"
	m["LogConfig.OutputPaths"] = "LOG_CONF_OUTPUT_PATH"
	m["LogConfig.ErrorOutputPaths"] = "LOG_CONF_ERROR_OUTPUT_PATH"

	return m
}
