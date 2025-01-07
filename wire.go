//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fabriciolfj/credit-card-service-go/configuration"
	"github.com/fabriciolfj/credit-card-service-go/kafka"
	"github.com/fabriciolfj/credit-card-service-go/listeners"
	"github.com/fabriciolfj/credit-card-service-go/producer"
	"github.com/fabriciolfj/credit-card-service-go/services"
	"github.com/google/wire"
)

func InitializeKafkaService() (*kafka.KafkaService, error) {
	wire.Build(
		configuration.ProvideKafkaProperties,
		configuration.ProvideKafkaConfig,
		listeners.ProvideCardApproveListener,
		producer.ProviderCardResultProducer,
		services.ProviderValidationService,
	)
	return &kafka.KafkaService{}, nil
}
