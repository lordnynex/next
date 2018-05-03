package mailers

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/sknv/upsale/app/core/cfg"
)

type Mailer struct {
	Auth smtp.Auth
	From string
}

func NewMailer() *Mailer {
	return &Mailer{
		Auth: smtp.PlainAuth(
			"", cfg.GetMailUsername(), cfg.GetMailPassword(), cfg.GetMailHost(),
		),
		From: "sail.notification@yandex.ru",
	}
}

func (m *Mailer) SendMail(subject, body, to string) error {
	// Log an email for the development mode.
	if !cfg.IsProduction() {
		log.Printf("info [send mail]: [%s] to [%s]", body, to)
		return nil
	}
	// Actually deliver an email.
	return smtp.SendMail(
		cfg.GetMailAddr(),
		m.Auth,
		m.From,
		[]string{to},
		[]byte(m.createMessage(subject, body, to)),
	)
}

func (m *Mailer) createMessage(subject, body, to string) string {
	return fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", m.From, to, subject, body)
}
