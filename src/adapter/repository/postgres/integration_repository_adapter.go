package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IntegrationRepositoryAdapter struct {
	*BaseRepository
}

func (i *IntegrationRepositoryAdapter) GetIntegrationByIdAndUserId(ctx context.Context, integrationId, userId int64) (*entity.IntegrationEntity, error) {
	var integrationModel model.IntegrationModel
	if err := i.db.WithContext(ctx).Model(&model.IntegrationModel{}).Where("id = ? AND user_id = ?", integrationId, userId).First(&integrationModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToIntegrationEntity(&integrationModel), nil
}

func (i *IntegrationRepositoryAdapter) GetIntegrationByUserId(ctx context.Context, userId int64) ([]*entity.IntegrationEntity, error) {
	var integrationModels []*model.IntegrationModel
	if err := i.db.WithContext(ctx).Model(&model.IntegrationModel{}).Where("user_id = ?", userId).Find(&integrationModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListIntegrationEntities(integrationModels), nil
}

func (i *IntegrationRepositoryAdapter) SaveIntegration(ctx context.Context, integration *entity.IntegrationEntity) (*entity.IntegrationEntity, error) {
	integrationModel := mapper.ToIntegrationModel(integration)
	if err := i.db.WithContext(ctx).Create(integrationModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToIntegrationEntity(integrationModel), nil
}

func NewIntegrationRepositoryAdapter(base *BaseRepository) port.IIntegrationPort {
	return &IntegrationRepositoryAdapter{
		base,
	}
}
