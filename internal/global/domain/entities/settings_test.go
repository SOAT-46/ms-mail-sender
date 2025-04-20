package entities_test

import (
	"testing"

	"github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
	t.Run("should create the settings successfully", func(t *testing.T) {
		// given & when
		settings := entities.NewSettings(
			"test@example.com",
			"smtp.example.com",
			587,
			"username",
			"password",
		)

		// then
		assert.Equal(t, "test@example.com", settings.From)
		assert.Equal(t, "smtp.example.com", settings.Host)
		assert.Equal(t, 587, settings.Port)
		assert.Equal(t, "username", settings.Username)
	})
}
