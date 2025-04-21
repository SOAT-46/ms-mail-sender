package messages

import "github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"

type SendMessage struct {
	To string `json:"to"`
}

func (msg *SendMessage) ToDomain(mailType entities.EmailType) entities.Email {
	return entities.Email{
		To:      msg.To,
		Subject: msg.getSubject(mailType),
		Type:    mailType,
	}
}

func (msg *SendMessage) getSubject(mailType entities.EmailType) string {
	if mailType == entities.Success {
		return "Video processado com sucesso"
	}
	return "Erro ao processar video"
}
