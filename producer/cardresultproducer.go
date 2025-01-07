package producer

import (
	"github.com/IBM/sarama"
	"github.com/fabriciolfj/credit-card-service-go/configuration"
)

type CardResultProducer struct {
	producer sarama.AsyncProducer
}

func ProviderCardResultProducer(cfg *configuration.KafkaConfig) (*CardResultProducer, error) {
	producer := &CardResultProducer{
		producer: cfg.Producer,
	}

	return producer, nil
}

func (p *CardResultProducer) Close() error {
	if p.producer != nil {
		return p.producer.Close()
	}
	return nil
}

func (p *CardResultProducer) SendMessage(message string) error {
	msg := &sarama.ProducerMessage{
		Topic: "seu-topico",
		Value: sarama.StringEncoder(message),
	}

	p.producer.Input() <- msg

	select {
	case err := <-p.producer.Errors():
		return err.Err
	case <-p.producer.Successes():
		return nil
	}
}
