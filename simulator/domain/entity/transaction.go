package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	id           string
	accountId    string
	amount       float64
	creditCard   CreditCard
	status       string
	errorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (transaction *Transaction) isValid() error {
	if transaction.amount > 1000 {
		return errors.New("you do not have limit for this transaction")
	}

	if transaction.amount < 1 {
		return errors.New("the amount must be greater than 1")
	}

	return nil
}

func (transaction *Transaction) setCreditCard(creditCard CreditCard) {
	transaction.creditCard = creditCard
}
