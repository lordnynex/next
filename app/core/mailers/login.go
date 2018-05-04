package mailers

import (
	"fmt"

	"github.com/jordan-wright/email"
)

type Login struct {
	*Mailer
}

func NewLogin() *Login {
	return &Login{Mailer: NewMailer()}
}

func (l *Login) Deliver(authSessionID, to string) {
	loginLink := "http://localhost:3000/login/" + authSessionID
	html := fmt.Sprintf(`<a href="%s">Log in</a>`, loginLink)
	text := "Paste this link into your web browser: " + loginLink
	email := &email.Email{
		To:      []string{to},
		Subject: "Login Link",
		HTML:    []byte(html),
		Text:    []byte(text),
	}

	l.Mailer.Deliver(email)
}
