package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToPipelineTemplateEntity(modelPT *model.PipelineTemplate) *entity.PipelineTemplateEntity {
	return &entity.PipelineTemplateEntity{
		BaseEntity:  entity.BaseEntity{ID: modelPT.ID},
		Name:        modelPT.Name,
		BuildTool:   modelPT.BuildTool,
		Description: modelPT.Description,
	}
}
