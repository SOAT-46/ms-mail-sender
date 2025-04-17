package services_test

import (
	"testing"

	entities2 "github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/services"
	"github.com/stretchr/testify/require"
)

func TestSendMailService(t *testing.T) {
	t.Run("should create the service successfully", func(t *testing.T) {
		// given
		settings := entities2.NewSettings(
			"test@example.com",
			"smtp.example.com",
			587,
			"username",
			"password",
		)

		// when
		service := services.NewSendMailService(settings, nil)

		// then
		require.NotNil(t, service, "should create the service")
	})
}
