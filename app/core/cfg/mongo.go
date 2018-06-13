package cfg

import (
	"os"
	"strings"
	"time"
)

const (
	// Authenticate against the "admin" database in a docker container.
	authDB              = "admin"
	defaultMongoTimeout = 60 * time.Second

	envMongoAddrs    = "NEXT_MONGO_ADDRS"
	envMongoDatabase = "MONGO_INITDB_DATABASE"
	envMongoUsername = "MONGO_INITDB_ROOT_USERNAME"
	envMongoPassword = "MONGO_INITDB_ROOT_PASSWORD"
	envMongoTimeout  = "NEXT_MONGO_TIMEOUT"
)

func GetMongoAddrs() []string {
	addrs, ok := os.LookupEnv(envMongoAddrs)
	if !ok {
		panic(envMongoAddrs + " is not defined")
	}
	return strings.Split(addrs, ",")
}

func GetMongoSource() string {
	return authDB
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
