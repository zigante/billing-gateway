package factory

import "github.com/zigante/billing-gateway/domain/repository"

type RepositoryFactory interface {
	createTransactionRepository() repository.TransactionRepository
}
