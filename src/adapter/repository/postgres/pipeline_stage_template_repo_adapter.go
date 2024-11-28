package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type PipelineStageTemplateRepoAdapter struct {
	*BaseRepository
}

func (p PipelineStageTemplateRepoAdapter) GetPipelineStageTemplateByTemplateID(ctx context.Context, templateID int64) ([]*entity.PipelineStageTemplateEntity, error) {
	var pipelineStageTemplateModel []*model.PipelineStageTemplateModel
	if err := p.db.WithContext(ctx).Model(&model.PipelineStageTemplateModel{}).Where("template_id = ?", templateID).Find(&pipelineStageTemplateModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToListPipelineStageTemplateEntity(pipelineStageTemplateModel), nil

}

func NewPipelineStageTemplateRepoAdapter(base *BaseRepository) port.IPipelineStageTemplatePort {
	return &PipelineStageTemplateRepoAdapter{
		BaseRepository: base,
	}
}
