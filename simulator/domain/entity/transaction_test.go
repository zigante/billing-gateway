package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionIsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.id = "1"
	transaction.accountId = "1"
	transaction.amount = 1000

	assert.Nil(t, transaction.isValid())
}

func TestTransactionIsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.id = "1"
	transaction.accountId = "1"
	transaction.amount = 1001

	err := transaction.isValid()
	assert.Error(t, err, "")
	assert.Equal(t, "you do not have limit for this transaction", err.Error())
}

func TestTransactionIsNotValidWithAmountLowerThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.id = "1"
	transaction.accountId = "1"
	transaction.amount = 0

	err := transaction.isValid()
	assert.Error(t, err, "")
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
