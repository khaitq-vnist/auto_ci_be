package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"gorm.io/gorm"
)

type UserRepositoryAdapter struct {
	BaseRepository
}

func (u UserRepositoryAdapter) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	userModel := &model.UserModel{}
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(userModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(userModel), nil
}

func (u UserRepositoryAdapter) SaveUser(ctx context.Context, tx *gorm.DB, userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	userModel := mapper.ToUserModel(userEntity)
	if err := tx.WithContext(ctx).Model(&model.UserModel{}).Create(userModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(userModel), nil
}

func NewUserRepositoryAdapter(db *gorm.DB) port.IUserPort {
	return &UserRepositoryAdapter{
		BaseRepository: BaseRepository{db: db},
	}
}
