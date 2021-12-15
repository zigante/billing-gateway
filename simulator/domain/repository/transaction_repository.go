package repository

type TransactionRepository interface {
	insert(id string, accountId string, amount float64, status string, errorMessage string) error
}
