package request

import (
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type BuddyCreateIntegrationRequest struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Scope string `json:"scope"`
	Token string `json:"token"`
}

func ToBuddyCreateIntegrationRequest(integration *entity.IntegrationEntity) *BuddyCreateIntegrationRequest {
	return &BuddyCreateIntegrationRequest{
		Name:  integration.Name,
		Type:  integration.Type,
		Scope: integration.Scope,
		Token: integration.AccessToken,
	}
}
