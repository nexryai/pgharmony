package config

import "os"

var (
	Port   = getEnvOrDefault("PORT", "3000")
	Secret = getEnvNotEmpty("PGHARMONY_SECRET")
)

func getEnvNotEmpty(env string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		panic("Environment variable " + env + " is not set")
	}
	return value
}

func getEnvOrDefault(env string, def string) string {
	if value, ok := os.LookupEnv(env); ok {
		return value
	}
	return def
}
