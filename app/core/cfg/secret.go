package cfg

import (
	"os"
)

const (
	defaultSecretKey = "XgK8SELWvdY7HEG9JAbjKgJj39RJJMyq"
)

func GetSecretKey() string {
	secretKey, ok := os.LookupEnv("APP_SECRET_KEY")
	if !ok {
		return defaultSecretKey
	}
	return secretKey
}
