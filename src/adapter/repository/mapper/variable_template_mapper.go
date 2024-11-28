package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToVariableTemplateEntity(variableTemplateModel *model.VariableTemplateModel) *entity.VariableTemplateEntity {
	return &entity.VariableTemplateEntity{
		BaseEntity: entity.BaseEntity{
			ID:        variableTemplateModel.ID,
			CreatedAt: variableTemplateModel.CreatedAt.Unix(),
			UpdatedAt: variableTemplateModel.UpdatedAt.Unix(),
		},
		StageID: variableTemplateModel.StageID,
		Key:     variableTemplateModel.Key,
		Value:   variableTemplateModel.Value,
	}
}
func ToListVariableTemplateEntity(variableTemplates []*model.VariableTemplateModel) []*entity.VariableTemplateEntity {
	var variableTemplateEntities []*entity.VariableTemplateEntity
	for _, variableTemplate := range variableTemplates {
		variableTemplateEntities = append(variableTemplateEntities, ToVariableTemplateEntity(variableTemplate))
	}
	return variableTemplateEntities
}
