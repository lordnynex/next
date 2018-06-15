package cfg

import (
	"os"
)

const (
	defaultEnv = "development"

	envEnv = "APP_ENV"
)

func IsProduction() bool {
	return getEnv() == "production"
}

func getEnv() string {
	env, ok := os.LookupEnv(envEnv)
	if !ok {
		return defaultEnv
	}
	return env
}
