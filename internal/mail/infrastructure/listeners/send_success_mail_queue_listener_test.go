package listeners_test

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners"
	"github.com/soat-46/ms-mail-sender/test/mail/domain/commands/doubles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errSendMailSuccessQueueListener = errors.New("failed to send error mail")

func TestSendSuccessMailQueueListener(t *testing.T) {
	t.Run("should log success when mail is sent successfully", func(t *testing.T) {
		// given
		hook := test.NewGlobal()
		defer hook.Reset()

		command := doubles.NewInMemorySendMailCommand()
		listener := listeners.NewSendSuccessMailQueueListener(command)

		// when
		listener.Run()

		// then
		require.Len(t, hook.Entries, 1, "should have one log entry")
		assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level, "should log at info level")
		assert.Equal(t, "success mail sent successfully", hook.LastEntry().Message, "should log success message")
	})

	t.Run("should log error when mail sending fails", func(t *testing.T) {
		// given
		hook := test.NewGlobal()
		defer hook.Reset()

		command := doubles.NewInMemorySendMailCommand().WithError(errSendMailSuccessQueueListener)
		listener := listeners.NewSendSuccessMailQueueListener(command)

		// when
		listener.Run()

		// then
		require.Len(t, hook.Entries, 1, "should have one log entry")
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level, "should log at error level")
		assert.Contains(t, hook.LastEntry().Message, "failed to send success mail", "should log error message")
		assert.Contains(t,
			hook.LastEntry().Message,
			errSendMailSuccessQueueListener.Error(),
			"should include error details in log")
	})
}
