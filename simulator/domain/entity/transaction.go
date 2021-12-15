package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	Id           string
	AccountId    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (transaction *Transaction) IsValid() error {
	if transaction.Amount > 1000 {
		return errors.New("you do not have limit for this transaction")
	}

	if transaction.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}

	return nil
}

func (transaction *Transaction) SetCreditCard(creditCard CreditCard) {
	transaction.CreditCard = creditCard
}
