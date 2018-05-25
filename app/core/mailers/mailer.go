package mailers

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/sknv/upsale/app/core/cfg"
	"github.com/sknv/upsale/app/core/initializers"
)

const (
	From = "Upsale Mailer <sail.notification@yandex.ru>"
)

type Mailer struct {
	From string
	Addr string
	Auth smtp.Auth
}

func NewMailer() *Mailer {
	return &Mailer{
		From: From,
		Addr: cfg.GetMailAddr(),
		Auth: smtp.PlainAuth(
			"", cfg.GetMailUsername(), cfg.GetMailPassword(), cfg.GetMailHost(),
		),
	}
}

func (m *Mailer) ExecuteTemplate(name string, data interface{}) []byte {
	bytes, err := initializers.GetHTML().ExecuteTemplateToBytes(name, data)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (m *Mailer) Deliver(email *email.Email) {
	// Log an email for the development mode.
	if !cfg.IsProduction() {
		log.Printf("[INFO] deliver email: %s to %s", email.Text, email.To)
		return
	}

	// Actually deliver an email.
	email.From = m.From
	if err := email.Send(m.Addr, m.Auth); err != nil {
		log.Print("[ERROR] deliver email: ", err)
	}
}
