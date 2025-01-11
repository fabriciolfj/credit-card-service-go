package services

import (
	"encoding/json"
	"fmt"
	"github.com/fabriciolfj/credit-card-service-go/client"
	"github.com/fabriciolfj/credit-card-service-go/entities"
	"github.com/fabriciolfj/credit-card-service-go/producer"
)

type ValidationService struct {
	producer *producer.CardResultProducer
	client   *client.RequestCard
}

func ProviderValidationService(producerMessage *producer.CardResultProducer, client *client.RequestCard) *ValidationService {
	return &ValidationService{producer: producerMessage, client: client}
}

func (vl *ValidationService) Execute(cardCustomer *entities.CardCustomer) error {
	result, err := vl.client.FindApprove(cardCustomer.Code)

	if err != nil {
		return fmt.Errorf("fail request %w, details %w", cardCustomer.Code, err.Error())
	}

	//todo send other topic
	if result == nil || result.Status != "APPROVED" {
		return fmt.Errorf("card denied %w", cardCustomer.Code)
	}

	message, err := json.Marshal(cardCustomer)
	if err != nil {
		return fmt.Errorf("fail serializer %w, details %w", cardCustomer.Code, err.Error())
	}

	if err := vl.producer.SendMessage(string(message)); err != nil {
		return fmt.Errorf("fail send message %w, details %w", cardCustomer.Code, err.Error())
	}

	return nil
}
