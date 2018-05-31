package cfg

import (
	"os"
)

func IsProduction() bool {
	return getEnv() == "production"
}

func getEnv() string {
	env, ok := os.LookupEnv("NEXT_ENV")
	if !ok {
		return "development" // Default application env.
	}
	return env
}
