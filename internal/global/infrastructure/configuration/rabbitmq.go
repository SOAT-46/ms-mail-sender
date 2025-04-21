package configuration

import (
	"fmt"
	"net"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
)

func OpenChannel(settings *entities.QueueSettings) (*amqp.Channel, error) {
	host := net.JoinHostPort(settings.Host, strconv.Itoa(settings.Port))
	url := fmt.Sprintf("amqp://%s:%s@%s/", settings.Username, settings.Password, host)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn.Channel()
}
