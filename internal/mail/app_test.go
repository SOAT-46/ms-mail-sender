package mail_test

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/soat-46/ms-mail-sender/internal/mail"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners"
	"github.com/soat-46/ms-mail-sender/test/mail/domain/commands/doubles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMailApp(t *testing.T) {
	t.Run("should create the app successfully", func(t *testing.T) {
		// given
		errorMailListener := &listeners.SendErrorMailQueueListener{}
		successMailListener := &listeners.SendSuccessMailQueueListener{}

		// when
		app := mail.NewApp(errorMailListener, successMailListener)

		// then
		require.NotNil(t, app, "should create the app")
	})

	t.Run("should create the listeners", func(t *testing.T) {
		// given
		hook := test.NewGlobal()
		defer hook.Reset()

		command := doubles.NewInMemorySendMailCommand()
		errorMailListener := listeners.NewSendErrorMailQueueListener(command)
		successMailListener := listeners.NewSendSuccessMailQueueListener(command)

		app := mail.NewApp(errorMailListener, successMailListener)

		// when
		app.RunConsumers()

		// then
		assert.Equal(t,
			"Queue consumers started successfully",
			hook.LastEntry().Message,
			"should log success message")
	})
}
