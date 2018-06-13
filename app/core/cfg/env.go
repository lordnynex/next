package cfg

import (
	"os"
)

const (
	envEnv = "APP_ENV"
)

func IsProduction() bool {
	return getEnv() == "production"
}

func getEnv() string {
	env, ok := os.LookupEnv(envEnv)
	if !ok {
		return "development" // Default application env.
	}
	return env
}
