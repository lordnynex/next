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

func (l *Login) Deliver(authSessionID, to string) {
	loginLink := "http://localhost:4000/login/" + authSessionID
	html := fmt.Sprintf(`<a href="%s">Log in</a>`, loginLink)
	text := "Paste this link into your web browser: " + loginLink
	if err := l.Send("Login Link", html, text, []string{to}); err != nil {
		log.Print("error [login deliver]: ", err)
	}
}
