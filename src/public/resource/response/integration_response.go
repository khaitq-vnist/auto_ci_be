package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type IntegrationResponse struct {
	ID               int64  `json:"id"`
	IntegrationName  string `json:"integration_name"`
	ProviderCode     string `json:"provider_code"`
	ProviderName     string `json:"provider_name"`
	ProviderUsername string `json:"provider_username"`
}

func ToIntegrationResponse(integration *entity.IntegrationEntity) *IntegrationResponse {
	return &IntegrationResponse{
		ID:               integration.ID,
		IntegrationName:  integration.Name,
		ProviderCode:     integration.ProviderCode,
		ProviderName:     integration.ProviderName,
		ProviderUsername: integration.ProviderUsername,
	}
}

func ToListIntegrationResponse(integrations []*entity.IntegrationEntity) []*IntegrationResponse {
	if integrations == nil {
		return nil
	}
	var integrationResponses []*IntegrationResponse
	for _, integration := range integrations {
		integrationResponses = append(integrationResponses, ToIntegrationResponse(integration))
	}
	return integrationResponses
}
