package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"

type BuddyCreateProjectResponse struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Url         string `json:"url"`
	Status      string `json:"status"`
}

func ToThirdPartyCreateProjectResponse(rsp *BuddyCreateProjectResponse) *response.ThirdPartyCreateProjectResponse {
	return &response.ThirdPartyCreateProjectResponse{
		Name:        rsp.Name,
		DisplayName: rsp.DisplayName,
		Url:         rsp.Url,
	}
}
