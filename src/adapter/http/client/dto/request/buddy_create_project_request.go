package request

import "github.com/khaitq-vnist/auto_ci_be/core/entity/dto/request"

type BuddyCreateProjectRequest struct {
	DisplayName       string                        `json:"display_name"`
	Name              string                        `json:"name,omitempty"`
	Integration       BuddyCreateProjectIntegration `json:"integration"`
	ExternalProjectID string                        `json:"external_project_id"`
}
type BuddyCreateProjectIntegration struct {
	HashID string `json:"hash_id"`
}

func ToBuddyCreateProjectRequest(dto *request.ThirdPartyCreateProjectRequest) *BuddyCreateProjectRequest {
	return &BuddyCreateProjectRequest{
		DisplayName: dto.DisplayName,
		Name:        dto.Name,
		Integration: BuddyCreateProjectIntegration{
			HashID: dto.Integration.HashID,
		},
		ExternalProjectID: dto.ExternalProjectID,
	}
}
