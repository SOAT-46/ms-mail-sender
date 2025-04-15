package doubles

import "github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"

type InMemoryMailSenderService struct {
	err error
}

func NewInMemoryMailSenderService() *InMemoryMailSenderService {
	return &InMemoryMailSenderService{}
}

func (service *InMemoryMailSenderService) WithOnError(err error) *InMemoryMailSenderService {
	service.err = err
	return service
}

func (service *InMemoryMailSenderService) Execute(_ entities.Email, _ string) error {
	return service.err
}
