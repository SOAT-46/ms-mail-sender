package listeners

import (
	logger "github.com/sirupsen/logrus"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
)

type SendErrorMailQueueListener struct {
	command commands.SendMail
}

func NewSendErrorMailQueueListener(command commands.SendMail) *SendErrorMailQueueListener {
	return &SendErrorMailQueueListener{command: command}
}

func (itself *SendErrorMailQueueListener) Run() {
	onSuccess := func() {
		logger.Info("error mail sent successfully")
	}
	onError := func(err error) {
		logger.Errorf("failed to send error mail. Reason: %v", err)
	}
	listeners := commands.SendMailListeners{
		OnSuccess: onSuccess,
		OnError:   onError,
	}
	itself.command.Execute(entities.Email{}, listeners)
}
