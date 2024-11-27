package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToPipelineStageTemplateEntity(pipelineStageTemplateModel *model.PipelineStageTemplateModel) *entity.PipelineStageTemplateEntity {
	return &entity.PipelineStageTemplateEntity{
		BaseEntity: entity.BaseEntity{
			ID:        pipelineStageTemplateModel.ID,
			CreatedAt: pipelineStageTemplateModel.CreatedAt.Unix(),
			UpdatedAt: pipelineStageTemplateModel.UpdatedAt.Unix(),
		},
		TemplateID: pipelineStageTemplateModel.TemplateID,
		StageID:    pipelineStageTemplateModel.StageID,
	}
}
