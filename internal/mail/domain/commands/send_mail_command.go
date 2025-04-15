package commands

import (
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/services"
)

type SendMailCommand struct {
	service services.MailSender
	render  services.RenderMail
}

func NewSendMailCommand(
	service services.MailSender,
	render services.RenderMail) *SendMailCommand {
	return &SendMailCommand{service: service, render: render}
}

func (cmd *SendMailCommand) Execute(entity entities.Email, listeners SendMailListeners) {
	body, err := cmd.render.Execute(entity.Type)
	if err != nil {
		listeners.OnError(err)
		return
	}
	err = cmd.service.Execute(entity, body)
	if err != nil {
		listeners.OnError(err)
	} else {
		listeners.OnSuccess()
	}
}
