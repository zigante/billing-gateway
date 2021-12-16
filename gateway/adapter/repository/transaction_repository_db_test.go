package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zigante/billing-gateway/adapter/repository/fixture"
	"github.com/zigante/billing-gateway/domain/entity"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")

	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")

	assert.Nil(t, err)
}
