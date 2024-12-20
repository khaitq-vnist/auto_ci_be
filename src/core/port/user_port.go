package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"gorm.io/gorm"
)

type IUserPort interface {
	SaveUser(ctx context.Context, tx *gorm.DB, userEntity *entity.UserEntity) (*entity.UserEntity, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}
