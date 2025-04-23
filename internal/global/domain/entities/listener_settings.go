package entities

type ListenerSettings struct {
	QueueURL string
	Topic    string
}

func NewListenerSettings(queueURL, topic string) *ListenerSettings {
	return &ListenerSettings{
		QueueURL: queueURL,
		Topic:    topic,
	}
}
