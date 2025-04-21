package listeners_test

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/global/infrastructure/configuration"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners/messages"
	"github.com/soat-46/ms-mail-sender/test/mail/domain/commands/doubles"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
)

type SendSuccessMailQueueListenerSuite struct {
	suite.Suite
	context   context.Context
	container *rabbitmq.RabbitMQContainer
	channel   *amqp.Channel
}

func (suite *SendSuccessMailQueueListenerSuite) SetupTest() {
	ctx := context.Background()
	suite.context = ctx

	container, err := rabbitmq.Run(
		ctx,
		RabbitMqImage,
		rabbitmq.WithAdminUsername("guest"),
		rabbitmq.WithAdminPassword("guest"))
	suite.Require().NoError(err)
	suite.container = container

	host, err := container.Host(ctx)
	suite.Require().NoError(err)

	port, err := suite.container.MappedPort(ctx, "5672")
	suite.Require().NoError(err)

	rabbitPort, err := strconv.Atoi(port.Port())
	suite.Require().NoError(err)

	username := suite.container.AdminUsername
	password := suite.container.AdminPassword

	settings := entities.NewQueueSettings(host, username, password, rabbitPort)
	channel, err := configuration.OpenChannel(settings)
	suite.Require().NoError(err)

	suite.channel = channel
}

func (suite *SendSuccessMailQueueListenerSuite) TearDownTest() {
	err := suite.container.Terminate(context.Background())
	suite.Require().NoError(err)
}

func (suite *SendSuccessMailQueueListenerSuite) TestSendErrorMailQueueListenerSuccess() {
	suite.Run("should send the success mail successfully", func() {
		// given
		hook := test.NewGlobal()
		defer hook.Reset()

		command := doubles.NewInMemorySendMailCommand()
		listener := listeners.NewSendSuccessMailQueueListener(command, suite.channel)

		message := messages.SendMessage{To: "test@example.com"}
		body, _ := json.Marshal(message)

		// when
		go listener.Run()

		time.Sleep(1 * time.Second)

		err := suite.channel.PublishWithContext(
			suite.context, "", "video-success", false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		suite.Require().NoError(err)

		time.Sleep(1 * time.Second)

		// then
		suite.Equal("Success mail sent successfully",
			hook.LastEntry().Message,
			"should log success message")
	})
}

func TestSendSuccessMailQueueListener(t *testing.T) {
	suite.Run(t, new(SendSuccessMailQueueListenerSuite))
}
