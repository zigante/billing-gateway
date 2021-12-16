package process_transaction

import (
	"github.com/zigante/billing-gateway/domain/entity"
	"github.com/zigante/billing-gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (processTransaction *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.Id = input.Id
	transaction.AccountId = input.AccountId
	transaction.Amount = input.Amount
	_, invalidCreditCardError := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	if invalidCreditCardError != nil {
		err := processTransaction.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, entity.REJECTED, invalidCreditCardError.Error())
		if err != nil {
			return TransactionDtoOutput{}, nil
		}

		output := TransactionDtoOutput{
			Id:           transaction.Id,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCreditCardError.Error(),
		}
		return output, nil
	}

	return TransactionDtoOutput{}, nil
}
