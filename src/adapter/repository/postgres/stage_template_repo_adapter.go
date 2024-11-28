package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type StageTemplateRepoAdapter struct {
	*BaseRepository
}

func (s StageTemplateRepoAdapter) GetStageTemplateByPipelineTemplateID(ctx context.Context, pipelineTemplateID int64) ([]*entity.StageTemplateEntity, error) {
	var stageTemplates []*model.StageTemplateModel
	if err := s.db.WithContext(ctx).Model(&model.StageTemplateModel{}).Where("pipeline_id = ?", pipelineTemplateID).Find(&stageTemplates).Error; err != nil {
		return nil, err
	}
	return mapper.ToListStageTemplateEntity(stageTemplates), nil
}

func (s StageTemplateRepoAdapter) GetStageTemplateByIds(ctx context.Context, IDs []int64) ([]*entity.StageTemplateEntity, error) {
	var stageTemplates []*model.StageTemplateModel
	if err := s.db.WithContext(ctx).Model(&model.StageTemplateModel{}).Where("id IN (?)", IDs).Find(&stageTemplates).Error; err != nil {
		return nil, err
	}
	return mapper.ToListStageTemplateEntity(stageTemplates), nil
}

func NewStageTemplateRepoAdapter(base *BaseRepository) port.IStageTemplatePort {
	return &StageTemplateRepoAdapter{
		BaseRepository: base,
	}
}
