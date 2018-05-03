package mailers

import (
	"fmt"
	"log"
)

type Login struct {
	*Mailer
}

func NewLogin() *Login {
	return &Login{Mailer: NewMailer()}
}

func (l *Login) Deliver(loginLink, to string) {
	body := fmt.Sprintf(
		"Paste this link into your web browser: http://localhost:4000/login/%s", loginLink,
	)
	if err := l.SendMail("Magic link", body, to); err != nil {
		log.Print("error [login deliver]: ", err)
	}
}
