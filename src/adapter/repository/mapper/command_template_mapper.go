package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToCommandTemplateEntity(commandTemplateModel *model.CommandTemplateModel) *entity.CommandTemplateEntity {
	return &entity.CommandTemplateEntity{
		BaseEntity: entity.BaseEntity{
			ID:        commandTemplateModel.ID,
			CreatedAt: commandTemplateModel.CreatedAt.Unix(),
			UpdatedAt: commandTemplateModel.UpdatedAt.Unix(),
		},
		Command: commandTemplateModel.Command,
		StageID: commandTemplateModel.StageId,
	}
}
func ToListCommandTemplateEntity(commandTemplates []*model.CommandTemplateModel) []*entity.CommandTemplateEntity {
	var commandTemplateEntities []*entity.CommandTemplateEntity
	for _, commandTemplate := range commandTemplates {
		commandTemplateEntities = append(commandTemplateEntities, ToCommandTemplateEntity(commandTemplate))
	}
	return commandTemplateEntities
}
