package cfg

import (
	"os"
)

func GetSecretKey() string {
	secretKey, ok := os.LookupEnv("APP_SECRET_KEY")
	if !ok {
		return "XgK8SELWvdY7HEG9JAbjKgJj39RJJMyq" // Default secret key.
	}
	return secretKey
}
