package mail

import (
	"github.com/google/wire"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/listeners"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/services"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	commands.Container,
	services.Container,
	listeners.Container,
	NewApp,
)
