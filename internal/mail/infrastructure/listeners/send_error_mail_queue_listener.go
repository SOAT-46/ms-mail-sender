package listeners

import (
	"encoding/json"
	"fmt"
	entities2 "github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/global/infrastructure/clients/contracts"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	logger "github.com/sirupsen/logrus"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners/messages"
)

type SendErrorMailQueueListener struct {
	command  commands.SendMail
	client   contracts.SQSClient
	settings *entities2.ListenerSettings
}

func NewSendErrorMailQueueListener(
	command commands.SendMail,
	client contracts.SQSClient,
	settings *entities2.ListenerSettings,
) *SendErrorMailQueueListener {
	return &SendErrorMailQueueListener{
		command:  command,
		client:   client,
		settings: settings,
	}
}

func (svc *SendErrorMailQueueListener) Run() {
	if svc.client == nil {
		logger.Errorf("client is nil")
		return
	}

	logger.Infof("Listening to queue: %s", svc.settings.Topic)

	fullURL := fmt.Sprintf("%s/%s", svc.settings.QueueURL, svc.settings.Topic)
	for {
		result, err := svc.client.ReceiveMessages(fullURL)
		if err != nil {
			logger.Errorf("Failed to receive messages. Reason: %v", err)
			break
		}

		for _, msg := range result.Messages {
			svc.process(fullURL, msg)
		}
	}
}

func (svc *SendErrorMailQueueListener) process(fullURL string, message types.Message) {
	var email messages.SendMessage
	logger.Info("Processing error mail")
	bytes := []byte(*message.Body)
	if err := json.Unmarshal(bytes, &email); err != nil {
		logger.Errorf("Failed to unmarshal message. Reason: %v", err)
		return
	}

	listeners := commands.SendMailListeners{
		OnSuccess: func() {
			logger.Info("Error mail sent successfully")
			err := svc.client.DeleteMessage(fullURL, message.ReceiptHandle)
			if err != nil {
				logger.Errorf("Failed to ack the message. Reason: %v", err)
			}
		},
		OnError: func(err error) {
			logger.Errorf("Failed to send error mail. Reason: %v", err)
		},
	}
	logger.Info("Sending error mail")
	svc.command.Execute(email.ToDomain(entities.Fail), listeners)
}
