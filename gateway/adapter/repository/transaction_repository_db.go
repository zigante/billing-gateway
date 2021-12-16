package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{
		db: db,
	}
}

func (transactionRepository *TransactionRepositoryDB) Insert(id string, accountId string, amount float64, status string, errorMessage string) error {
	statement, err := transactionRepository.db.Prepare(`INSERT INTO TRANSACTIONS VALUES ($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(id, accountId, amount, status, errorMessage, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
