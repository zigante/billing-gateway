package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zigante/billing-gateway/adapter/broker/kafka"
	"github.com/zigante/billing-gateway/adapter/factory"
	"github.com/zigante/billing-gateway/adapter/presenter/transaction"
	"github.com/zigante/billing-gateway/usecase/process_transaction"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	producerConfigMap := &ckafka.ConfigMap{"bootstrap.servers": "host.docker.internal:9094"}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewProducer(producerConfigMap, kafkaPresenter)

	messageChannel := make(chan *ckafka.Message)
	consumerConfigMap := &ckafka.ConfigMap{"bootstrap.servers": "host.docker.internal:9094", "client.id": "goapp", "group.id": "goapp"}
	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(consumerConfigMap, topics)
	go consumer.Consume(messageChannel)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for message := range messageChannel {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(message.Value, &input)
		usecase.Execute(input)
	}
}
