package services

import (
	"github.com/google/wire"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/services"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewRenderMailTemplate,
	wire.Bind(new(services.RenderMail), new(*RenderMailTemplate)),
	NewSendMailService,
	wire.Bind(new(services.MailSender), new(*SendMailService)),
)
