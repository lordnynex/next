package cfg

import (
	"os"
	"strings"
	"time"
)

const (
	defaultMongoTimeout = 60 * time.Second

	envMongoAddrs    = "UPSALE_MONGO_ADDRS"
	envMongoDatabase = "UPSALE_MONGO_DATABASE"
	envMongoUsername = "UPSALE_MONGO_USERNAME"
	envMongoPassword = "UPSALE_MONGO_PASSWORD"
	envMongoTimeout  = "UPSALE_MONGO_TIMEOUT"
)

func GetMongoAddrs() []string {
	addrs, ok := os.LookupEnv(envMongoAddrs)
	if !ok {
		panic(envMongoAddrs + " is not defined")
	}
	return strings.Split(addrs, ",")
}

func GetMongoDatabase() string {
	database, ok := os.LookupEnv(envMongoDatabase)
	if !ok {
		panic(envMongoDatabase + " is not defined")
	}
	return database
}

func GetMongoUsername() string {
	username, ok := os.LookupEnv(envMongoUsername)
	if !ok {
		panic(envMongoUsername + " is not defined")
	}
	return username
}

func GetMongoPassword() string {
	password, ok := os.LookupEnv(envMongoPassword)
	if !ok {
		panic(envMongoPassword + " is not defined")
	}
	return password
}

func GetMongoTimeout() time.Duration {
	duration, ok := os.LookupEnv(envMongoTimeout)
	if !ok {
		return defaultMongoTimeout
	}

	timeout, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return timeout
}
