package doubles

import "github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"

type InMemoryRenderMailService struct {
	err error
}

func NewInMemoryRenderMailService() *InMemoryRenderMailService {
	return &InMemoryRenderMailService{}
}

func (service *InMemoryRenderMailService) WithOnError(err error) *InMemoryRenderMailService {
	service.err = err
	return service
}

func (service *InMemoryRenderMailService) Execute(_ entities.EmailType) (string, error) {
	return "", service.err
}
