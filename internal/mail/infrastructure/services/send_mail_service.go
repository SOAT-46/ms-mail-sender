package services

import (
	"errors"
	"fmt"

	entities2 "github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"gopkg.in/gomail.v2"
)

var (
	ErrSendMail = errors.New("failed to send email")
)

type SendMailService struct {
	from string
	mail *gomail.Dialer
}

func NewSendMailService(
	settings *entities2.Settings,
	mail *gomail.Dialer) *SendMailService {
	return &SendMailService{
		from: settings.From,
		mail: mail,
	}
}

func (itself *SendMailService) Execute(mail entities.Email, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", itself.from)
	message.SetHeader("To", mail.To)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody("text/html", body)

	err := itself.mail.DialAndSend(message)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrSendMail, err)
	}
	return nil
}
