package entities

type Settings struct {
	From     string
	Host     string
	Port     int
	Username string
	Password string
}

func NewSettings(from, host string, port int, username, password string) *Settings {
	return &Settings{from, host, port, username, password}
}
