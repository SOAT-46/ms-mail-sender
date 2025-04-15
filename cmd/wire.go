//go:build !test && wireinject

package main

import (
	"os"
	"strconv"

	"github.com/google/wire"
	"github.com/soat-46/ms-mail-sender/internal/global/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail"
	"gopkg.in/gomail.v2"
)

func injectApps() []entities.App {
	wire.Build(
		injectSettings,
		injectGoMail,
		mail.Container,
		newApps,
	)
	return nil
}

func injectSettings() *entities.Settings {
	from := os.Getenv("MAIL_FROM")
	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	return entities.NewSettings(from, host, port, username, password)
}

func injectGoMail(settings *entities.Settings) *gomail.Dialer {
	return gomail.NewDialer(settings.Host, settings.Port, settings.Username, settings.Password)
}

func newApps(mail *mail.App) []entities.App {
	return []entities.App{mail}
}
