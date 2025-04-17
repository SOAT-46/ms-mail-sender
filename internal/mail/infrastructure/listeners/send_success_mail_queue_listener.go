package listeners

import (
	logger "github.com/sirupsen/logrus"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
)

type SendSuccessMailQueueListener struct {
	command commands.SendMail
}

func NewSendSuccessMailQueueListener(command commands.SendMail) *SendSuccessMailQueueListener {
	return &SendSuccessMailQueueListener{command: command}
}

func (itself *SendSuccessMailQueueListener) Run() {
	onSuccess := func() {
		logger.Info("success mail sent successfully")
	}
	onError := func(err error) {
		logger.Errorf("failed to send success mail. Reason: %v", err)
	}
	listeners := commands.SendMailListeners{
		OnSuccess: onSuccess,
		OnError:   onError,
	}
	itself.command.Execute(entities.Email{}, listeners)
}
