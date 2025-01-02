package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToExecutionHistoryModel(entity *entity.ExecutionHistoryEntity) *model.ExecutionHistoryModel {
	return &model.ExecutionHistoryModel{
		BaseModel:         model.BaseModel{ID: entity.ID},
		ProjectID:         entity.ProjectID,
		PipelineID:        entity.PipelineID,
		ThirdPartyID:      entity.ThirdPartyID,
		ThirdPartyProject: entity.ThirdPartyProject,
		LogsFile:          entity.LogsFile,
		Coverage:          entity.Coverage,
	}
}
func ToExecutionHistoryEntity(model *model.ExecutionHistoryModel) *entity.ExecutionHistoryEntity {
	return &entity.ExecutionHistoryEntity{
		BaseEntity:        entity.BaseEntity{ID: model.ID},
		ProjectID:         model.ProjectID,
		PipelineID:        model.PipelineID,
		ThirdPartyID:      model.ThirdPartyID,
		ThirdPartyProject: model.ThirdPartyProject,
		LogsFile:          model.LogsFile,
		Coverage:          model.Coverage,
	}
}
func ToListExecutionHistoryEntity(models []*model.ExecutionHistoryModel) []*entity.ExecutionHistoryEntity {
	var entities []*entity.ExecutionHistoryEntity
	for _, model := range models {
		entities = append(entities, ToExecutionHistoryEntity(model))
	}
	return entities
}
func ToListExecutionHistoryModel(entities []*entity.ExecutionHistoryEntity) []*model.ExecutionHistoryModel {
	var models []*model.ExecutionHistoryModel
	for _, entity := range entities {
		models = append(models, ToExecutionHistoryModel(entity))
	}
	return models
}
