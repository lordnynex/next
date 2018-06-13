package cfg

import (
	"fmt"
	"os"
)

const (
	MailBaseUrl = "http://localhost:8080" // TODO: define a production domain.

	envMailHost     = "APP_MAIL_HOST"
	envMailPort     = "APP_MAIL_PORT"
	envMailUsername = "APP_MAIL_USERNAME"
	envMailPassword = "APP_MAIL_PASSWORD"
)

func GetMailAddr() string {
	return fmt.Sprintf("%s:%s", GetMailHost(), GetMailPort())
}

func GetMailHost() string {
	host, ok := os.LookupEnv(envMailHost)
	if !ok {
		panic(envMailHost + " is not defined")
	}
	return host
}

func GetMailPort() string {
	port, ok := os.LookupEnv(envMailPort)
	if !ok {
		panic(envMailPort + " is not defined")
	}
	return port
}

func GetMailUsername() string {
	username, ok := os.LookupEnv(envMailUsername)
	if !ok {
		panic(envMailUsername + " is not defined")
	}
	return username
}

func GetMailPassword() string {
	pass, ok := os.LookupEnv(envMailPassword)
	if !ok {
		panic(envMailPassword + " is not defined")
	}
	return pass
}
