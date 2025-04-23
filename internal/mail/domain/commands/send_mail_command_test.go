package commands_test

import (
	"errors"
	"testing"

	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"github.com/soat-46/ms-mail-sender/test/mail/infrastructure/services/doubles"
	"github.com/stretchr/testify/assert"
)

var (
	errSendMailCommand = errors.New("failed to send mail")
	errRenderTemplate  = errors.New("failed to render template")
)

func TestSendMailCommand(t *testing.T) {
	email := entities.Email{
		To:      "test@example.com",
		Subject: "Test Subject",
		Type:    entities.Success,
	}
	t.Run("should call OnSuccess when mail is rendered and sent successfully", func(t *testing.T) {
		// given
		var isSuccess = false
		mailService := doubles.NewInMemoryMailSenderService()
		renderService := doubles.NewInMemoryRenderMailService()
		command := commands.NewSendMailCommand(mailService, renderService)

		listeners := commands.SendMailListeners{
			OnSuccess: func() { isSuccess = true },
			OnError:   func(_ error) { assert.Fail(t, "should not call OnError") },
		}

		// when
		command.Execute(email, listeners)

		// then
		assert.True(t, isSuccess)
	})

	t.Run("should call OnError when template rendering fails", func(t *testing.T) {
		// given
		mailService := doubles.NewInMemoryMailSenderService()
		renderService := doubles.NewInMemoryRenderMailService().WithOnError(errRenderTemplate)
		command := commands.NewSendMailCommand(mailService, renderService)

		var capturedError error
		listeners := commands.SendMailListeners{
			OnSuccess: func() { assert.Fail(t, "should not call OnSuccess") },
			OnError: func(err error) {
				capturedError = err
			},
		}

		// when
		command.Execute(email, listeners)

		// then
		assert.ErrorIs(t, capturedError, errRenderTemplate, "should call OnError when template rendering fails")
	})

	t.Run("should call OnError when mail sending fails", func(t *testing.T) {
		// given
		mailService := doubles.NewInMemoryMailSenderService().WithOnError(errSendMailCommand)
		renderService := doubles.NewInMemoryRenderMailService()
		command := commands.NewSendMailCommand(mailService, renderService)

		var capturedError error
		listeners := commands.SendMailListeners{
			OnSuccess: func() { assert.Fail(t, "should not call OnSuccess") },
			OnError:   func(err error) { capturedError = err },
		}

		// when
		command.Execute(email, listeners)

		// then
		assert.ErrorIs(t, capturedError, errSendMailCommand, "should call OnError when mail sending fails")
	})
}
