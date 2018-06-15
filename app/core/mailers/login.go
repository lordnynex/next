package mailers

import (
	"fmt"

	"github.com/go-chi/render"
	"github.com/jordan-wright/email"

	"github.com/sknv/next/app/core/cfg"
	"github.com/sknv/next/app/core/models"
)

type Login struct {
	*Mailer
}

func NewLogin() *Login {
	return &Login{Mailer: NewMailer()}
}

func (l *Login) Deliver(user *models.User) {
	loginLink := fmt.Sprintf(
		"%s/auth/login?email=%s&password=%s", cfg.MailBaseUrl, user.Email, user.Code,
	)
	plain := fmt.Sprintf(
		`Paste this link into web browser to login into your accout: %s
Or just type the disposable code right in the app: %s`, loginLink, user.Code,
	)

	email := &email.Email{
		To:      []string{user.Email},
		Subject: "Login Link",
		Text:    []byte(plain),
		HTML: l.Mailer.ExecuteTemplate(
			"login",
			render.M{"loginLink": loginLink, "code": user.Code},
		),
	}
	l.Mailer.Deliver(email)
}
