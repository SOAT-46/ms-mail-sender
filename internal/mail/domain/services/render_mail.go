package services

import "github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"

type RenderMail interface {
	Execute(mailType entities.EmailType) (string, error)
}
