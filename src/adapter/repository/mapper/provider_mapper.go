package mapper

import (
	model2 "github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToProviderModel(providerEntity *entity.ProviderEntity) *model2.ProviderModel {
	if providerEntity == nil {
		return nil
	}
	return &model2.ProviderModel{
		BaseModel: model2.BaseModel{
			ID: providerEntity.ID,
		},
		Name:    providerEntity.Name,
		Code:    providerEntity.Code,
		ApiUrl:  providerEntity.ApiUrl,
		AuthUrl: providerEntity.AuthUrl,
	}
}

func ToProviderEntity(providerModel *model2.ProviderModel) *entity.ProviderEntity {
	if providerModel == nil {
		return nil
	}
	return &entity.ProviderEntity{
		BaseEntity: entity.BaseEntity{
			ID:        providerModel.ID,
			CreatedAt: providerModel.CreatedAt.Unix(),
			UpdatedAt: providerModel.UpdatedAt.Unix(),
		},
		Name:    providerModel.Name,
		Code:    providerModel.Code,
		ApiUrl:  providerModel.ApiUrl,
		AuthUrl: providerModel.AuthUrl,
	}
}
