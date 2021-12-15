package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCredicCardNumber(t *testing.T) {
	_, err := NewCreditCard("12345", "Random Name For Testing", 12, 2021, 123)
	assert.Equal(t, "invalid credit card number", err.Error())
}

func TestCreditCardExpiredMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Random Name For Testing", 13, 2021, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Random Name For Testing", 0, 2021, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Random Name For Testing", 7, 2021, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("4193523830170205", "Random Name For Testing", 7, lastYear.Year(), 123)

	assert.Equal(t, "invalid expiration year", err.Error())
}
