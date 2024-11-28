package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type PipelineTemplateRepositoryAdapter struct {
	*BaseRepository
}

func (p PipelineTemplateRepositoryAdapter) GetPipelineTemplateByBuildTools(ctx context.Context, typePipeline string) (*entity.PipelineTemplateEntity, error) {
	var pipelineTemplateModel model.PipelineTemplateModel
	if err := p.db.WithContext(ctx).Model(&model.PipelineTemplateModel{}).Where("build_tool = ?", typePipeline).First(&pipelineTemplateModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToPipelineTemplateEntity(&pipelineTemplateModel), nil
}

func NewPipelineTemplateRepositoryAdapter(base *BaseRepository) port.IPipelineTemplatePort {
	return &PipelineTemplateRepositoryAdapter{
		BaseRepository: base,
	}
}
