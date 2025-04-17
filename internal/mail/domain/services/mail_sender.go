package services

import "github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"

type MailSender interface {
	Execute(mail entities.Email, body string) error
}
