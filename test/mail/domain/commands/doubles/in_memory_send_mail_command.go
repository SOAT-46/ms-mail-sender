package doubles

import (
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/commands"
	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
)

type InMemorySendMailCommand struct {
	err error
}

func NewInMemorySendMailCommand() *InMemorySendMailCommand {
	return &InMemorySendMailCommand{}
}

func (cmd *InMemorySendMailCommand) WithError(err error) *InMemorySendMailCommand {
	cmd.err = err
	return cmd
}

func (cmd *InMemorySendMailCommand) Execute(_ entities.Email, listeners commands.SendMailListeners) {
	if cmd.err != nil {
		listeners.OnError(cmd.err)
		return
	}
	listeners.OnSuccess()
}
