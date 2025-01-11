//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fabriciolfj/credit-card-service-go/client"
	"github.com/fabriciolfj/credit-card-service-go/configuration"
	"github.com/fabriciolfj/credit-card-service-go/listeners"
	"github.com/fabriciolfj/credit-card-service-go/producer"
	"github.com/fabriciolfj/credit-card-service-go/services"
	"github.com/google/wire"
)

func InitializeApp() (*listeners.CardApproveListener, error) {
	wire.Build(
		configuration.ProvideKafkaProperties,
		configuration.ProvideKafkaConfig,
		producer.ProviderCardResultProducer,
		client.ProvideRequestCard,
		services.ProviderValidationService,
		listeners.ProvideCardApproveListener,
	)
	return &listeners.CardApproveListener{}, nil
}
