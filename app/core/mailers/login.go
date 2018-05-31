package mailers

import (
	"fmt"

	"github.com/go-chi/render"
	"github.com/jordan-wright/email"

	"github.com/sknv/next/app/core/cfg"
)

type Login struct {
	*Mailer
}

func NewLogin() *Login {
	return &Login{Mailer: NewMailer()}
}

func (l *Login) Deliver(authSessionID, to string) {
	loginLink := fmt.Sprintf("%s/login/%s", cfg.MailBaseUrl, authSessionID)
	email := &email.Email{
		To:      []string{to},
		Subject: "Login Link",
		HTML:    l.Mailer.ExecuteTemplate("login", render.M{"link": loginLink}),
		Text: []byte(
			"Paste this link into web browser to login into your accout: " + loginLink,
		),
	}
	l.Mailer.Deliver(email)
}
