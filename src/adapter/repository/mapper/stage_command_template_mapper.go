package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToStageCommandTemplateEntity(stageCommandTemplateModel *model.StageCommandTemplateModel) *entity.StageCommandTemplateEntity {
	return &entity.StageCommandTemplateEntity{
		BaseEntity: entity.BaseEntity{
			ID:        stageCommandTemplateModel.ID,
			CreatedAt: stageCommandTemplateModel.CreatedAt.Unix(),
			UpdatedAt: stageCommandTemplateModel.UpdatedAt.Unix(),
		},
		StageID:   stageCommandTemplateModel.StageID,
		CommandID: stageCommandTemplateModel.CommandID,
	}
}
