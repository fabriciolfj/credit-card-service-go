package listeners

import (
	"github.com/IBM/sarama"
	"github.com/fabriciolfj/credit-card-service-go/configuration"
	"github.com/fabriciolfj/credit-card-service-go/services"
)

type MessageHandler func(message string)

type CardApproveListener struct {
	consumer sarama.Consumer
	service  services.ValidationService
}

func ProvideCardApproveListener(cfg *configuration.KafkaConfig, service services.ValidationService) (*CardApproveListener, error) {
	listener := &CardApproveListener{
		consumer: cfg.Consumer,
		service:  service,
	}
	return listener, nil
}
