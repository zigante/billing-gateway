package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionIsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 1000

	assert.Nil(t, transaction.IsValid())
}

func TestTransactionIsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 1001

	err := transaction.IsValid()
	assert.Error(t, err, "")
	assert.Equal(t, "you do not have limit for this transaction", err.Error())
}

func TestTransactionIsNotValidWithAmountLowerThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 0

	err := transaction.IsValid()
	assert.Error(t, err, "")
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
