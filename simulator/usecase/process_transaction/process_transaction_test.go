package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zigante/billing-gateway/domain/entity"
	mock_repository "github.com/zigante/billing-gateway/domain/repository/mock"
)

func TestProcessTransactionExecuteValidTransaction(t *testing.T) {
	input := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "Random name for testing",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    700,
	}
	expectedOutput := TransactionDtoOutput{
		Id:     "1",
		Status: entity.APPROVED,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransactionExecuteInvalidTransaction(t *testing.T) {
	input := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "Random name for testing",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    1200,
	}
	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you do not have limit for this transaction",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransactionExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "12345",
		CreditCardName:            "Random name for testing",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}
	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
