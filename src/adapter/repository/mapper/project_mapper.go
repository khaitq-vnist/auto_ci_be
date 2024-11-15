package mapper

import (
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToProjectModel(project *entity.ProjectEntity) *model.ProjectModel {
	if project == nil {
		return nil
	}
	return &model.ProjectModel{
		BaseModel: model.BaseModel{
			ID: project.ID,
		},
		Name:        project.Name,
		FullName:    project.FullName,
		Private:     project.Private,
		OwnerID:     project.OwnerID,
		Description: project.Description,
	}
}
func ToProjectEntity(project *model.ProjectModel) *entity.ProjectEntity {
	if project == nil {
		return nil
	}
	return &entity.ProjectEntity{
		BaseEntity: entity.BaseEntity{
			ID:        project.ID,
			CreatedAt: project.CreatedAt.Unix(),
			UpdatedAt: project.UpdatedAt.Unix(),
		},
		Name:        project.Name,
		FullName:    project.FullName,
		Private:     project.Private,
		OwnerID:     project.OwnerID,
		Description: project.Description,
	}
}
