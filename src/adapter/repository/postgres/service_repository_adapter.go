package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"gorm.io/gorm"
)

type ServiceRepositoryAdapter struct {
	BaseRepository
}

func (s ServiceRepositoryAdapter) GetAllService(ctx context.Context) ([]*entity.ServiceEntity, error) {
	var serviceModels []*model.ServiceModel
	if err := s.db.Model(&model.ServiceModel{}).Find(&serviceModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListServiceEntityMapper(serviceModels), nil
}

func NewServiceRepositoryAdapter(db *gorm.DB) port.IServicePort {
	return &ServiceRepositoryAdapter{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
}
