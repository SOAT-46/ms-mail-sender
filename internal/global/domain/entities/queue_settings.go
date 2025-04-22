package entities

type QueueSettings struct {
	Host     string
	Username string
	Password string
	Port     int
}

func NewQueueSettings(host, username, password string, port int) *QueueSettings {
	return &QueueSettings{
		Host:     host,
		Password: password,
		Username: username,
		Port:     port,
	}
}
