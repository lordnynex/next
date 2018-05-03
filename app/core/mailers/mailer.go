package mailers

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/sknv/upsale/app/core/cfg"
)

type Mailer struct {
	Addr string
	Auth smtp.Auth
	From string
}

func NewMailer() *Mailer {
	return &Mailer{
		// TODO: Specify for production.
		From: "Upsale Mailer <sail.notification@yandex.ru>",

		Addr: cfg.GetMailAddr(),
		Auth: smtp.PlainAuth(
			"", cfg.GetMailUsername(), cfg.GetMailPassword(), cfg.GetMailHost(),
		),
	}
}

func (m *Mailer) Send(subject, html, text string, to []string) error {
	// Log an email for the development mode.
	if !cfg.IsProduction() {
		log.Printf("info [send email]: %s to %s", text, to)
		return nil
	}

	// Actually deliver an email.
	email := &email.Email{
		From:    m.From,
		To:      to,
		Subject: subject,
		HTML:    []byte(html),
		Text:    []byte(text),
	}
	return email.Send(m.Addr, m.Auth)
}
