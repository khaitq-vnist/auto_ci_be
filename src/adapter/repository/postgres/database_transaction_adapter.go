package postgres

import (
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"gorm.io/gorm"
)

type DatabaseTransactionAdapter struct {
	*BaseRepository
}

func (d *DatabaseTransactionAdapter) StartTransaction() *gorm.DB {
	return d.BaseRepository.StartTransaction()
}
func (d *DatabaseTransactionAdapter) Commit(txDB *gorm.DB) error {
	return d.BaseRepository.CommitTransaction(txDB)
}

func (d *DatabaseTransactionAdapter) Rollback(txDB *gorm.DB) error {
	return d.BaseRepository.RollbackTransaction(txDB)
}

func NewDatabaseTransactionAdapter(db *gorm.DB) port.IDatabaseTransactionPort {
	return &DatabaseTransactionAdapter{
		BaseRepository: &BaseRepository{db},
	}
}
