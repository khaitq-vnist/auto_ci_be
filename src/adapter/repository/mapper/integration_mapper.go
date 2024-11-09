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
	}
}
