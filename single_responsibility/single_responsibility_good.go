package single_responsibility

import (
	"fmt"
	"log"
	"net/smtp"
)

type EmailGorm struct {
	GormModel
	From    string
	To      string
	Subject string
	Message string
}

type EmailRepository interface {
	Save(from string, to string, subject string, message string) error
}

type EmailDBRepository struct {
	db *Gorm
}

func NewEmailRepository(db *Gorm) EmailRepository {
	return &EmailDBRepository{
		db: db,
	}
}

func (r *EmailDBRepository) Save(from string, to string, subject string, message string) error {
	email := EmailGorm{
		From:    from,
		To:      to,
		Subject: subject,
		Message: message,
	}

	err := r.db.Create(&email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type EmailSender interface {
	Send(from string, to string, subject string, message string) error
}

type EmailSMTPSender struct {
	smtpHost     string
	smtpPassword string
	smtpPort     int
}

func NewEmailSender(smtpHost string, smtpPassword string, smtpPort int) EmailSender {
	return &EmailSMTPSender{
		smtpHost:     smtpHost,
		smtpPassword: smtpPassword,
		smtpPort:     smtpPort,
	}
}

func (s *EmailSMTPSender) Send(from string, to string, subject string, message string) error {
	auth := smtp.PlainAuth("", from, s.smtpPassword, s.smtpHost)

	server := fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort)

	err := smtp.SendMail(server, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type EmailServiceGood struct {
	repository EmailRepository
	sender     EmailSender
}

func NewEmailServiceGood(repository EmailRepository, sender EmailSender) *EmailServiceGood {
	return &EmailServiceGood{
		repository: repository,
		sender:     sender,
	}
}

func (s *EmailServiceGood) Send(from string, to string, subject string, message string) error {
	err := s.repository.Save(from, to, subject, message)
	if err != nil {
		return err
	}

	return s.sender.Send(from, to, subject, message)
}
