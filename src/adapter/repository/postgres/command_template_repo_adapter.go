package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type CommandTemplateRepoAdapter struct {
	*BaseRepository
}

func (c CommandTemplateRepoAdapter) GetCommandTemplateByStageIDs(ctx context.Context, ids []int64) ([]*entity.CommandTemplateEntity, error) {
	var models []*model.CommandTemplateModel
	if err := c.db.WithContext(ctx).Where("stage_id IN (?)", ids).Find(&models).Error; err != nil {
		return nil, err
	}
	return mapper.ToListCommandTemplateEntity(models), nil
}

func NewCommandTemplateRepoAdapter(base *BaseRepository) port.ICommandTemplatePort {
	return &CommandTemplateRepoAdapter{
		base,
	}
}
