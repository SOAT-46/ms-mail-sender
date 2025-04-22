package listeners

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners/messages"
)

const (
	errorQueue = "video-error"
)

type SendErrorMailQueueListener struct {
	command commands.SendMail
	channel *amqp.Channel
}

func NewSendErrorMailQueueListener(command commands.SendMail, channel *amqp.Channel) *SendErrorMailQueueListener {
	return &SendErrorMailQueueListener{command: command, channel: channel}
}

func (svc *SendErrorMailQueueListener) Run() {
	if svc.channel == nil {
		logger.Errorf("Channel is nil")
		return
	}

	_, err := svc.channel.QueueDeclare(
		errorQueue,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logger.Fatalf("Failed to declare queue: %v", err)
	}

	err = svc.channel.Qos(1, 0, false)
	if err != nil {
		logger.Fatalf("Failed to set QoS: %v", err)
	}

	msgs, err := svc.channel.Consume(
		errorQueue,
		"fail-mail-sender",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to register consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			svc.process(msg)
		}
	}()

	logger.Infof("Listening to queue: %s", errorQueue)
	select {} // Keep the goroutine alive
}

func (svc *SendErrorMailQueueListener) process(message amqp.Delivery) {
	var email messages.SendMessage
	logger.Info("Processing error mail")
	if err := json.Unmarshal(message.Body, &email); err != nil {
		logger.Errorf("Failed to unmarshal message. Reason: %v", err)
		err = message.Nack(false, false) // Don't requeue bad messages
		if err != nil {
			logger.Errorf("Failed to requeue error messages. Reason: %v", err)
		}
		return
	}

	listeners := commands.SendMailListeners{
		OnSuccess: func() {
			logger.Info("Error mail sent successfully")
			err := message.Ack(false)
			if err != nil {
				logger.Errorf("Failed to ack the message. Reason: %v", err)
			}
		},
		OnError: func(err error) {
			logger.Errorf("Failed to send error mail. Reason: %v", err)
			nackErr := message.Nack(false, true) // Requeue for retry
			if nackErr != nil {
				logger.Errorf("Failed to nack the message. Reason: %v", nackErr)
			}
		},
	}
	logger.Info("Sending error mail")
	svc.command.Execute(email.ToDomain(entities.Fail), listeners)
}
