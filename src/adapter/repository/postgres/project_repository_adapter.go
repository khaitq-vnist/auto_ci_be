package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ProjectRepositoryAdapter struct {
	*BaseRepository
}

func (p ProjectRepositoryAdapter) GetListProjectByUserId(ctx context.Context, userId int64) ([]*entity.ProjectEntity, error) {
	var projectModels []*model.ProjectModel
	if err := p.db.WithContext(ctx).Model(&model.ProjectModel{}).Where("owner_id = ?", userId).Find(&projectModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListProjectEntities(projectModels), nil
}

func (p ProjectRepositoryAdapter) SaveProject(ctx context.Context, entity *entity.ProjectEntity) (*entity.ProjectEntity, error) {
	projectModel := mapper.ToProjectModel(entity)
	if err := p.db.WithContext(ctx).Create(&projectModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToProjectEntity(projectModel), nil
}

func NewProjectRepositoryAdapter(base *BaseRepository) port.IProjectPort {
	return &ProjectRepositoryAdapter{
		BaseRepository: base,
	}
}
