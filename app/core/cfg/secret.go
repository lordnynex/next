package cfg

import "os"

const (
	envSecretKey = "UPSALE_SECRET_KEY"
)

func GetSecretKey() string {
	secretKey, ok := os.LookupEnv(envSecretKey)
	if !ok {
		panic(envSecretKey + " is not defined")
	}
	return secretKey
}
