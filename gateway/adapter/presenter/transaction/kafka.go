package transaction

import (
	"encoding/json"

	"github.com/zigante/billing-gateway/usecase/process_transaction"
)

type KafkaPresenter struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (transaction *KafkaPresenter) Bind(input interface{}) error {
	transaction.Id = input.(process_transaction.TransactionDtoOutput).Id
	transaction.Status = input.(process_transaction.TransactionDtoOutput).Status
	transaction.ErrorMessage = input.(process_transaction.TransactionDtoOutput).ErrorMessage

	return nil
}

func (transaction *KafkaPresenter) Show() ([]byte, error) {
	outputJSON, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}

	return outputJSON, nil
}
