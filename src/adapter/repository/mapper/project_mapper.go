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
		Name:                project.Name,
		FullName:            project.FullName,
		Private:             project.Private,
		OwnerID:             project.OwnerID,
		Description:         project.Description,
		IntegrationId:       project.IntegrationId,
		ThirdPartyProjectID: project.ThirdPartyProjectID,
		SonarProjectName:    project.SonarProjectName,
		SonarKey:            project.SonarKey,
		SonarToken:          project.SonarToken,
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
		Name:                project.Name,
		FullName:            project.FullName,
		Private:             project.Private,
		OwnerID:             project.OwnerID,
		Description:         project.Description,
		IntegrationId:       project.IntegrationId,
		ThirdPartyProjectID: project.ThirdPartyProjectID,
		SonarProjectName:    project.SonarProjectName,
		SonarKey:            project.SonarKey,
		SonarToken:          project.SonarToken,
	}
}

func ToListProjectEntities(projectModels []*model.ProjectModel) []*entity.ProjectEntity {
	var projectEntities []*entity.ProjectEntity
	for _, projectModel := range projectModels {
		projectEntities = append(projectEntities, ToProjectEntity(projectModel))
	}
	return projectEntities
}
