package listeners

import (
	"github.com/google/wire"
	entities2 "github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/global/infrastructure/clients/contracts"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"os"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	ProvideSuccessListener,
	ProvideErrorListener,
)

func ProvideSuccessListener(
	command commands.SendMail,
	client contracts.SQSClient) *SendSuccessMailQueueListener {
	settings := &entities2.ListenerSettings{
		QueueURL: os.Getenv("AWS_SQS_URL"),
		Topic:    os.Getenv("AWS_SUCCESS_TOPIC"),
	}
	return NewSendSuccessMailQueueListener(command, client, settings)
}

func ProvideErrorListener(
	command commands.SendMail,
	client contracts.SQSClient) *SendErrorMailQueueListener {
	settings := &entities2.ListenerSettings{
		QueueURL: os.Getenv("AWS_SQS_URL"),
		Topic:    os.Getenv("AWS_ERROR_TOPIC"),
	}
	return NewSendErrorMailQueueListener(command, client, settings)
}
