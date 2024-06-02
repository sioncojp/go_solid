package single_responsibility

import (
	"fmt"
	"log"
	"net/smtp"
)

type Email struct {
	From    string
	To      string
	Subject string
	Message string
}

type EmailServiceBad struct {
	db           Gorm
	smtpHost     string
	smtpPassword string
	smtpPort     int
}

func NewEmailServiceBad(smtpHost string, smtpPassword string, smtpPort int) *EmailServiceBad {
	return &EmailServiceBad{
		smtpHost:     smtpHost,
		smtpPassword: smtpPassword,
		smtpPort:     smtpPort,
	}
}

func (s *EmailServiceBad) Send(from string, to string, subject string, message string) error {
	email := Email{
		From:    from,
		To:      to,
		Subject: subject,
		Message: message,
	}

	err := s.db.Create(&email)
	if err != nil {
		log.Println(err)
		return err
	}

	auth := smtp.PlainAuth("", from, s.smtpPassword, s.smtpHost)

	server := fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort)

	err = smtp.SendMail(server, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
