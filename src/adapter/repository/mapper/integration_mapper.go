package mapper

import (
	model2 "github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToIntegrationModel(integrationEntity *entity.IntegrationEntity) *model2.IntegrationModel {
	if integrationEntity == nil {
		return nil
	}
	return &model2.IntegrationModel{
		BaseModel: model2.BaseModel{
			ID: integrationEntity.ID,
		},
		Name:             integrationEntity.Name,
		UserId:           integrationEntity.UserId,
		ProviderId:       integrationEntity.ProviderId,
		ProviderName:     integrationEntity.ProviderName,
		AccessToken:      integrationEntity.AccessToken,
		ProviderUsername: integrationEntity.ProviderUsername,
		ThirdPartyHashId: integrationEntity.ThirdPartyHashId,
	}
}
func ToIntegrationEntity(integrationModel *model2.IntegrationModel) *entity.IntegrationEntity {
	if integrationModel == nil {
		return nil
	}
	return &entity.IntegrationEntity{
		BaseEntity: entity.BaseEntity{
			ID: integrationModel.ID,
		},
		Name:             integrationModel.Name,
		UserId:           integrationModel.UserId,
		ProviderId:       integrationModel.ProviderId,
		ProviderName:     integrationModel.ProviderName,
		AccessToken:      integrationModel.AccessToken,
		ProviderUsername: integrationModel.ProviderUsername,
		ThirdPartyHashId: integrationModel.ThirdPartyHashId,
	}
}
func ToListIntegrationEntities(integrationModels []*model2.IntegrationModel) []*entity.IntegrationEntity {
	if integrationModels == nil {
		return nil
	}
	var integrationEntities []*entity.IntegrationEntity
	for _, integrationModel := range integrationModels {
		integrationEntities = append(integrationEntities, ToIntegrationEntity(integrationModel))
	}
	return integrationEntities
}
