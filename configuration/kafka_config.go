package configuration

import (
	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	Producer sarama.AsyncProducer
	Consumer sarama.ConsumerGroup
}

func ProvideKafkaConfig(cfg *KafkaProperties) (*KafkaConfig, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRoundRobin(),
	}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	producer, err := sarama.NewAsyncProducer([]string{cfg.BootstrapServers}, config)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumerGroup([]string{cfg.BootstrapServers}, cfg.GroupID, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConfig{
		Producer: producer,
		Consumer: consumer,
	}, nil
}
