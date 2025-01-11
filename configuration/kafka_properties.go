package configuration

import (
	"github.com/magiconair/properties"
)

type KafkaProperties struct {
	BootstrapServers string
	GroupID          string
	ConsumerTopic    string
	ProducerTopic    string
	AutoOffsetReset  string
}

func ProvideKafkaProperties() (*KafkaProperties, error) {
	p, err := properties.LoadFile("config.properties", properties.UTF8)
	if err != nil {
		return nil, err
	}

	return &KafkaProperties{
		BootstrapServers: p.GetString("kafka.bootstrap.servers", ""),
		GroupID:          p.GetString("kafka.group.id", ""),
		ConsumerTopic:    p.GetString("kafka.topic.consumer", ""),
		ProducerTopic:    p.GetString("kafka.topic.producer", ""),
		AutoOffsetReset:  p.GetString("kafka.auto.offset.reset", "earliest"),
	}, nil
}
