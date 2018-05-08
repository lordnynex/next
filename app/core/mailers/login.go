package mailers

import (
	"fmt"

	"github.com/jordan-wright/email"

	"github.com/sknv/upsale/app/core/cfg"
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
		HTML:    []byte(l.getHTML(loginLink)),
		Text:    []byte(l.getText(loginLink)),
	}

	l.Mailer.Deliver(email)
}

func (l *Login) getHTML(loginLink string) string {
	html := fmt.Sprintf(`<p>
Click the link below to login into your account.
This link will expire in 15 minutes and can only be used once.
</p>
<a href="%s">Log in</a>`,
		loginLink,
	)
	return html
}

func (l *Login) getText(loginLink string) string {
	return "Paste this link into your web browser to login: " + loginLink
}
