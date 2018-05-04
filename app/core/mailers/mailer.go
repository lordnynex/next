package mailers

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/sknv/upsale/app/core/cfg"
)

type Mailer struct {
	From string
	Addr string
	Auth smtp.Auth
}

func NewMailer() *Mailer {
	return &Mailer{
		From: "Upsale Mailer <sail.notification@yandex.ru>",

		Addr: cfg.GetMailAddr(),
		Auth: smtp.PlainAuth(
			"", cfg.GetMailUsername(), cfg.GetMailPassword(), cfg.GetMailHost(),
		),
	}
}

func (m *Mailer) Deliver(email *email.Email) {
	// Log an email for the development mode.
	if !cfg.IsProduction() {
		log.Printf("info [deliver email]: %s to %s", email.Text, email.To)
		return
	}

	// Actually deliver an email.
	email.From = m.From
	if err := email.Send(m.Addr, m.Auth); err != nil {
		log.Print("error [deliver email]: ", err)
	}
}
