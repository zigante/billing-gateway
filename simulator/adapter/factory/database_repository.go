package factory

import (
	"database/sql"

	repositoryAdapter "github.com/zigante/billing-gateway/adapter/repository"
	"github.com/zigante/billing-gateway/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repositoryAdapter.NewTransactionRepositoryDb(r.DB)
}
