package services

import (
	"github.com/fabriciolfj/credit-card-service-go/entities"
	"github.com/fabriciolfj/credit-card-service-go/producer"
)

type ValidationService struct {
	producer *producer.CardResultProducer
}

func ProviderValidationService(producerMessage *producer.CardResultProducer) *ValidationService {
	return &ValidationService{producer: producerMessage}
}

func (vl *ValidationService) Execute(cardCustomer *entities.CardCustomer) error {
	err := vl.producer.SendMessage("")
	if err != nil {
		return err
	}

	return nil
}
