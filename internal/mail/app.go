package mail

import (
	logger "github.com/sirupsen/logrus"
	"github.com/soat-46/ms-mail-sender/internal/global/domain/queues"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners"
)

type App struct {
	listeners []queues.Listener
}

func NewApp(
	sendErrorMail *listeners.SendErrorMailQueueListener,
	sendSuccessMail *listeners.SendSuccessMailQueueListener,
) *App {
	consumers := []queues.Listener{
		sendErrorMail,
		sendSuccessMail,
	}
	return &App{consumers}
}

func (itself *App) RunConsumers() {
	logger.Info("Starting queue consumers...")
	for _, consumer := range itself.listeners {
		consumer.Run()
	}
	logger.Info("Queue consumers started successfully")
}
