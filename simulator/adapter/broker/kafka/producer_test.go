package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/zigante/billing-gateway/adapter/presenter/transaction"
	"github.com/zigante/billing-gateway/domain/entity"
	"github.com/zigante/billing-gateway/usecase/process_transaction"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		Id:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you do not have limit for this transaction",
	}

	configMap := ckafka.ConfigMap{"test.mock.num.brokers": 3}
	producer := NewProducer(&configMap, transaction.NewTransactionKafkaPresenter())

	err := producer.Publish(expectedOutput, []byte("1"), "test")

	assert.Nil(t, err)

}
