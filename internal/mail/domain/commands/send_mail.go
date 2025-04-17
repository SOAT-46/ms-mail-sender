package commands

import "github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"

type SendMail interface {
	Execute(entity entities.Email, listeners SendMailListeners)
}

type SendMailListeners struct {
	OnSuccess func()
	OnError   func(err error)
}
