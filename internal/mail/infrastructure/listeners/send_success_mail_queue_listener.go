package listeners

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	entities2 "github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/global/infrastructure/clients/contracts"

	logger "github.com/sirupsen/logrus"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners/messages"
)

type SendSuccessMailQueueListener struct {
	command  commands.SendMail
	client   contracts.SQSClient
	settings *entities2.ListenerSettings
}

func NewSendSuccessMailQueueListener(
	command commands.SendMail,
	client contracts.SQSClient,
	settings *entities2.ListenerSettings,
) *SendSuccessMailQueueListener {
	return &SendSuccessMailQueueListener{
		command:  command,
		client:   client,
		settings: settings,
	}
}

func (svc *SendSuccessMailQueueListener) Run() {
	if svc.client == nil {
		logger.Errorf("Client is nil")
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

func (svc *SendSuccessMailQueueListener) process(fullURL string, message types.Message) {
	var email messages.SendMessage
	logger.Info("Processing success mail")
	bytes := []byte(*message.Body)
	if err := json.Unmarshal(bytes, &email); err != nil {
		logger.Errorf("Failed to unmarshal message. Reason: %v", err)
		return
	}

	onSuccess := func() {
		logger.Info("Success mail sent successfully")
		err := svc.client.DeleteMessage(fullURL, message.ReceiptHandle)
		if err != nil {
			logger.Errorf("Failed to ack the message. Reason: %v", err)
		}
	}
	onError := func(err error) {
		logger.Errorf("failed to send success mail. Reason: %v", err)
	}
	listeners := commands.SendMailListeners{
		OnSuccess: onSuccess,
		OnError:   onError,
	}
	svc.command.Execute(email.ToDomain(entities.Success), listeners)
}
