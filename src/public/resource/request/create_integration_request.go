package request

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type CreateIntegrationRequest struct {
	IntegrationName string `json:"integration_name"`
	ProviderCode    string `json:"provider_code"`
	AccessToken     string `json:"access_token"`
}

func (r *CreateIntegrationRequest) ToEntity() *entity.IntegrationEntity {
	return &entity.IntegrationEntity{
		Name:         r.IntegrationName,
		ProviderCode: r.ProviderCode,
		AccessToken:  r.AccessToken,
	}
}
