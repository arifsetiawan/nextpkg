package mailer

import (
	"crypto/tls"

	gomail "gopkg.in/gomail.v2"
)

// Mailer ...
type Mailer struct {
	Host               string
	Port               int
	Username           string
	Password           string
	InsecureSkipVerify bool
}

// MailData ...
type MailData struct {
	From    string
	Tos     []string
	Subject string
	Body    string
}

// Send ...
func (m *Mailer) Send(data *MailData) (err error) {

	e := gomail.NewMessage()
	e.SetHeader("From", data.From)
	e.SetHeader("To", data.Tos...)
	e.SetHeader("Subject", data.Subject)
	e.SetBody("text/html", data.Body)

	d := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: m.InsecureSkipVerify}

	if err := d.DialAndSend(e); err != nil {
		return err
	}

	return
}
