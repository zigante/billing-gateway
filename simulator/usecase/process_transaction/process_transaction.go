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
	creditCard, invalidCreditCardError := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	if invalidCreditCardError != nil {
		return processTransaction.rejectTransaction(transaction, invalidCreditCardError)
	}

	transaction.SetCreditCard(*creditCard)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return processTransaction.rejectTransaction(transaction, invalidTransaction)
	}

	return processTransaction.approveTransaction(input, transaction)
}

func (processTransaction *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := processTransaction.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, entity.REJECTED, invalidTransaction.Error())
	if err != nil {
		panic(err)
	}
	output := TransactionDtoOutput{
		Id:           transaction.Id,
		Status:       entity.REJECTED,
		ErrorMessage: invalidTransaction.Error(),
	}

	return output, nil
}

func (processTransaction *ProcessTransaction) approveTransaction(input TransactionDtoInput, transaction *entity.Transaction) (TransactionDtoOutput, error) {
	err := processTransaction.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, entity.APPROVED, "")
	if err != nil {
		panic(err)
	}
	output := TransactionDtoOutput{
		Id:           transaction.Id,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	return output, nil
}
