package services_test

import (
	"testing"

	smtpmock "github.com/mocktools/go-smtp-mock"
	entities2 "github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/services"
	"github.com/stretchr/testify/suite"
	"gopkg.in/gomail.v2"
)

type SendMailServiceSuite struct {
	suite.Suite
	server *smtpmock.Server
}

func (suite *SendMailServiceSuite) SetupTest() {
	suite.server = smtpmock.New(smtpmock.ConfigurationAttr{
		//HostAddress:       "smtp.example.com",
		//PortNumber:        2525,
		LogToStdout:       true,
		LogServerActivity: true,
	})
}

func (suite *SendMailServiceSuite) TearDownTest() {
	err := suite.server.Stop()
	suite.Require().NoError(err)
}

func (suite *SendMailServiceSuite) TestSendMailService() {
	suite.Run("should send the mail successfully", func() {
		// given
		initServer := suite.server.Start()
		suite.Require().NoError(initServer)

		hostAddress, portNumber := "127.0.0.1", suite.server.PortNumber

		settings := entities2.NewSettings(
			"test@example.com",
			hostAddress,
			portNumber,
			"",
			"",
		)
		dialer := gomail.NewDialer(settings.Host, settings.Port, settings.Username, settings.Password)
		service := services.NewSendMailService(settings, dialer)

		email := entities.Email{
			To:      "test@example.com",
			Subject: "Test Subject",
			Type:    entities.Success,
		}

		// when
		err := service.Execute(email, "")

		// then
		suite.NoError(err, "should send the mail")
	})

	suite.Run("should fail to send the mail", func() {
		// given
		settings := entities2.NewSettings(
			"test@example.com",
			"127.0.0.1",
			2525,
			"",
			"",
		)
		dialer := gomail.NewDialer(settings.Host, settings.Port, settings.Username, settings.Password)
		service := services.NewSendMailService(settings, dialer)

		email := entities.Email{
			To:      "test@example.com",
			Subject: "Test Subject",
			Type:    entities.Fail,
		}

		// when
		err := service.Execute(email, "")

		// then
		suite.Error(err, "should fail to send the mail")
	})
}

func TestSendMailServiceSuite(t *testing.T) {
	suite.Run(t, new(SendMailServiceSuite))
}
