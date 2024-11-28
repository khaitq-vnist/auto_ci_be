package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToStageTemplateEntity(stageTemplateModel *model.StageTemplateModel) *entity.StageTemplateEntity {
	return &entity.StageTemplateEntity{
		BaseEntity: entity.BaseEntity{
			ID:        stageTemplateModel.ID,
			CreatedAt: stageTemplateModel.CreatedAt.Unix(),
			UpdatedAt: stageTemplateModel.UpdatedAt.Unix(),
		},
		PipelineTemplateID: stageTemplateModel.PipelineTemplateID,
		Name:               stageTemplateModel.Name,
		Type:               stageTemplateModel.Type,
		DockerImage:        stageTemplateModel.DockerImage,
		DockerImageTag:     stageTemplateModel.DockerImageTag,
	}
}
func ToListStageTemplateEntity(stageTemplates []*model.StageTemplateModel) []*entity.StageTemplateEntity {
	var stageTemplateEntities []*entity.StageTemplateEntity
	for _, stageTemplate := range stageTemplates {
		stageTemplateEntities = append(stageTemplateEntities, ToStageTemplateEntity(stageTemplate))
	}
	return stageTemplateEntities
}
