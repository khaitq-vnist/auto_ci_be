package postgres

import "gorm.io/gorm"

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) StartTransaction() *gorm.DB {
	return r.db.Begin()
}
func (r *BaseRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}
func (r *BaseRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}
