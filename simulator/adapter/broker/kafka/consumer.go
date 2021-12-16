package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{ConfigMap: configMap, Topics: topics}
}

func (consumer *Consumer) Consume(channel chan *ckafka.Message) error {
	_consumer, err := ckafka.NewConsumer(consumer.ConfigMap)
	if err != nil {
		return err
	}

	err = _consumer.SubscribeTopics(consumer.Topics, nil)
	if err != nil {
		return err
	}

	for {
		message, err := _consumer.ReadMessage(-1)
		if err == nil {
			channel <- message
		}
	}
}
