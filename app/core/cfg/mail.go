package cfg

import (
	"fmt"
	"os"
)

const (
	MailBaseUrl = "http://localhost:3000"

	envMailHost     = "NEXT_MAIL_HOST"
	envMailPort     = "NEXT_MAIL_PORT"
	envMailUsername = "NEXT_MAIL_USERNAME"
	envMailPassword = "NEXT_MAIL_PASSWORD"
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
