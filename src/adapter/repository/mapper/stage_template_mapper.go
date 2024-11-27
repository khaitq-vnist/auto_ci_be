package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToStageTemplateEntity(stageTemplateModel *model.StageTemplate) *entity.StageTemplateEntity {
	return &entity.StageTemplateEntity{
		BaseEntity: entity.BaseEntity{
			ID:        stageTemplateModel.ID,
			CreatedAt: stageTemplateModel.CreatedAt.Unix(),
			UpdatedAt: stageTemplateModel.UpdatedAt.Unix(),
		},
		Name:           stageTemplateModel.Name,
		Type:           stageTemplateModel.Type,
		DockerImage:    stageTemplateModel.DockerImage,
		DockerImageTag: stageTemplateModel.DockerImageTag,
	}
}
