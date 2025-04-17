package commands

import "github.com/google/wire"

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewSendMailCommand,
	wire.Bind(new(SendMail), new(*SendMailCommand)),
)
