package listeners

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/fabriciolfj/credit-card-service-go/configuration"
	"github.com/fabriciolfj/credit-card-service-go/entities"
	"github.com/fabriciolfj/credit-card-service-go/services"
	properties "github.com/magiconair/properties"
	"log"
)

type MessageHandler func(message string)

type CardApproveListener struct {
	consumer sarama.ConsumerGroup
	service  *services.ValidationService
	topic    string
}

func ProvideCardApproveListener(cfg *configuration.KafkaConfig, validation *services.ValidationService) (*CardApproveListener, error) {
	prop, error := properties.LoadFile("config.properties", properties.UTF8)

	if error != nil {
		panic(error)
	}

	listener := &CardApproveListener{
		consumer: cfg.Consumer,
		service:  validation,
		topic:    prop.GetString("topic.request.approve", ""),
	}
	return listener, nil
}

func (c *CardApproveListener) Start() error {
	ctx := context.Background()

	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		err := c.consumer.Consume(ctx, []string{c.topic}, c)
		if err != nil {
			return fmt.Errorf("erro ao consumir mensagem: %w", err)
		}
	}
}

func (c *CardApproveListener) Setup(sarama.ConsumerGroupSession) error {
	log.Println("init consumer group...")
	return nil
}

func (c *CardApproveListener) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("finish consumer group...")
	return nil
}

func (c *CardApproveListener) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := c.handleMessage(message); err != nil {
			log.Printf("Erro ao processar mensagem: %v", err)
		}

		session.MarkMessage(message, "")
	}
	return nil
}

func (c *CardApproveListener) Close() error {
	if err := c.consumer.Close(); err != nil {
		return fmt.Errorf("erro close consumer group: %w", err)
	}
	log.Println("close consumer group...")
	return nil
}

func (c *CardApproveListener) handleMessage(message *sarama.ConsumerMessage) error {
	log.Printf("message received - Tópico: %s, Partição: %d, Offset: %d",
		message.Topic, message.Partition, message.Offset)

	var cardCustomer entities.CardCustomer
	if err := json.Unmarshal(message.Value, &cardCustomer); err != nil {
		return fmt.Errorf("error deserializer message: %w", err)
	}

	if err := c.service.Execute(&cardCustomer); err != nil {
		return fmt.Errorf("error processo message: %w", err)
	}

	log.Printf("message process success - ID: %s", cardCustomer.Code)
	return nil
}
