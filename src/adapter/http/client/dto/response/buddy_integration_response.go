package response

import (
	response2 "github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
)

type BuddyIntegrationResponse struct {
	HashID string `json:"hash_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Scope  string `json:"scope"`
}

func ToIntegrationResponseDto(it *BuddyIntegrationResponse) *response2.ThirdPartyCreateIntegrationResponse {
	return &response2.ThirdPartyCreateIntegrationResponse{
		HashID: it.HashID,
		Name:   it.Name,
		Type:   it.Type,
		Scope:  it.Scope,
	}
}
