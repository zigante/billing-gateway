package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/zigante/billing-gateway/adapter/presenter"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
}

func NewkafkaProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter) *Producer {
	return &Producer{
		ConfigMap: configMap,
		Presenter: presenter,
	}
}

func (producer *Producer) Publish(message interface{}, key []byte, topic string) error {
	_producer, err := ckafka.NewProducer(producer.ConfigMap)
	if err != nil {
		return err
	}

	err = producer.Presenter.Bind(message)
	if err != nil {
		return err
	}

	parsedMessage, err := producer.Presenter.Show()
	if err != nil {
		return err
	}

	kafkaMessage := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          parsedMessage,
		Key:            key,
	}

	err = _producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}

	return nil
}
