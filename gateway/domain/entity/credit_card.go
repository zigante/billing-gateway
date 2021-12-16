package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	Number          string
	Name            string
	ExpirationMonth int
	ExpirationYear  int
	CVV             int
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, cvv int) (*CreditCard, error) {
	creditCard := &CreditCard{
		Number:          number,
		Name:            name,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		CVV:             cvv,
	}

	err := creditCard.IsValid()
	if err != nil {
		return nil, err
	}

	return creditCard, nil
}

func (creditCard *CreditCard) IsValid() error {
	err := creditCard.validateNumber()

	if err != nil {
		return err
	}

	err = creditCard.validateMonth()

	if err != nil {
		return err
	}

	err = creditCard.validateYear()

	if err != nil {
		return err
	}

	return nil
}

func (creditCard *CreditCard) validateNumber() error {
	regex := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	if !regex.MatchString(creditCard.Number) {
		return errors.New("invalid credit card number")
	}

	return nil
}

func (creditCard *CreditCard) validateMonth() error {
	if creditCard.ExpirationMonth > 0 && creditCard.ExpirationMonth < 13 {
		return nil
	}

	return errors.New("invalid expiration month")
}

func (creditCard *CreditCard) validateYear() error {
	if creditCard.ExpirationYear >= time.Now().Year() {
		return nil
	}

	return errors.New("invalid expiration year")
}
