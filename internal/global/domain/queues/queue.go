package queues

type Message struct {
	Key   string
	Value []byte
}

type Queue interface {
	Publish(topic string, msg Message) error
	Consume(topic string, handler func(Message)) error
	Close() error
}
