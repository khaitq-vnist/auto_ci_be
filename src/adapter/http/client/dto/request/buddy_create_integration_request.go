package request

import (
	"github.com/khaitq-vnist/auto_ci_be/core/constant"
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
		Type:  constant.GITHUB_INTEGRATION,
		Scope: constant.WORKSPACE_SCOPE,
		Token: integration.AccessToken,
	}
}
