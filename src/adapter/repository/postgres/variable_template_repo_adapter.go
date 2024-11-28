package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type VariableTemplateRepoAdapter struct {
	*BaseRepository
}

func (v VariableTemplateRepoAdapter) GetVariableTemplateByStageIDs(ctx context.Context, stageIDs []int64) ([]*entity.VariableTemplateEntity, error) {
	var models []*model.VariableTemplateModel
	if err := v.db.WithContext(ctx).Where("stage_id IN (?)", stageIDs).Find(&models).Error; err != nil {
		return nil, err
	}
	return mapper.ToListVariableTemplateEntity(models), nil
}

func NewVariableTemplateRepoAdapter(base *BaseRepository) port.IVariableTemplatePort {
	return &VariableTemplateRepoAdapter{
		base,
	}
}
