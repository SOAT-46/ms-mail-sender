package entities_test

import (
	"testing"

	"github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewQueueSettings(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		username string
		password string
		port     int
		want     *entities.QueueSettings
	}{
		{
			name:     "normal settings",
			host:     "rabbitmq.example.com",
			username: "admin",
			password: "securepassword",
			port:     5672,
			want: &entities.QueueSettings{
				Host:     "rabbitmq.example.com",
				Username: "admin",
				Password: "securepassword",
				Port:     5672,
			},
		},
		{
			name:     "empty credentials",
			host:     "localhost",
			username: "",
			password: "",
			port:     5672,
			want: &entities.QueueSettings{
				Host:     "localhost",
				Username: "",
				Password: "",
				Port:     5672,
			},
		},
		{
			name:     "special characters in credentials",
			host:     "mq.example.net",
			username: "user@domain",
			password: "p@$$w0rd!",
			port:     5671,
			want: &entities.QueueSettings{
				Host:     "mq.example.net",
				Username: "user@domain",
				Password: "p@$$w0rd!",
				Port:     5671,
			},
		},
		{
			name:     "minimum port number",
			host:     "127.0.0.1",
			username: "guest",
			password: "guest",
			port:     1,
			want: &entities.QueueSettings{
				Host:     "127.0.0.1",
				Username: "guest",
				Password: "guest",
				Port:     1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := entities.NewQueueSettings(tt.host, tt.username, tt.password, tt.port)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestQueueSettingsFields(t *testing.T) {
	t.Run("verify all fields are set correctly", func(t *testing.T) {
		qs := entities.NewQueueSettings("broker.example.com", "mqadmin", "s3cr3t", 5672)

		assert.Equal(t, "broker.example.com", qs.Host)
		assert.Equal(t, "mqadmin", qs.Username)
		assert.Equal(t, "s3cr3t", qs.Password)
		assert.Equal(t, 5672, qs.Port)
	})

	t.Run("verify field types", func(t *testing.T) {
		qs := entities.NewQueueSettings("", "", "", 0)

		assert.IsType(t, "", qs.Host)
		assert.IsType(t, "", qs.Username)
		assert.IsType(t, "", qs.Password)
		assert.IsType(t, 0, qs.Port)
	})
}

func TestQueueSettingsEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		port int
	}{
		{"minimum valid port", 1},
		{"maximum valid port", 65535},
		{"zero port", 0},
		{"negative port", -1},
		{"default RabbitMQ port", 5672},
		{"RabbitMQ TLS port", 5671},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := entities.NewQueueSettings("", "", "", tt.port)
			assert.Equal(t, tt.port, qs.Port)
		})
	}
}

func TestQueueSettingsPointer(t *testing.T) {
	t.Run("constructor returns pointer", func(t *testing.T) {
		qs := entities.NewQueueSettings("", "", "", 0)
		assert.IsType(t, &entities.QueueSettings{}, qs)
	})

	t.Run("pointer is not nil", func(t *testing.T) {
		qs := entities.NewQueueSettings("", "", "", 0)
		assert.NotNil(t, qs)
	})
}

func TestQueueSettingsFieldOrder(t *testing.T) {
	t.Run("verify field order in constructor", func(t *testing.T) {
		// This test ensures the constructor parameters match the struct field order
		qs := entities.NewQueueSettings("host", "user", "pass", 1234)

		assert.Equal(t, "host", qs.Host)
		assert.Equal(t, "user", qs.Username)
		assert.Equal(t, "pass", qs.Password)
		assert.Equal(t, 1234, qs.Port)
	})
}
