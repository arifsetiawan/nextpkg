package mailer

import (
	"os"
	"strconv"
	"testing"
)

func TestSendEmail(t *testing.T) {

	port, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))

	mailer := &Mailer{
		Host:               os.Getenv("MAILER_HOST"),
		Port:               port,
		Username:           os.Getenv("MAILER_AUTH_USER"),
		Password:           os.Getenv("MAILER_AUTH_PASS"),
		InsecureSkipVerify: true,
	}

	mailData := &MailData{
		From:    "one@nextflow.tech",
		Tos:     []string{"arif.setiawan@notmymail.com"},
		Subject: "Test",
		Body:    "Hello",
	}

	err := mailer.Send(mailData)
	if err != nil {
		t.Error(err)
	}
}
