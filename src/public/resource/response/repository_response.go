package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity/dto"

type ReposResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Private   bool   `json:"private"`
	FullName  string `json:"full_name"`
	CreateAt  int64  `json:"create_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func ToListReposResponse(repos []*dto.ThirdPartyProviderReposResponse) []*ReposResponse {
	var response []*ReposResponse
	for _, repo := range repos {
		response = append(response, &ReposResponse{
			ID:        repo.ID,
			Name:      repo.Name,
			Private:   repo.Private,
			FullName:  repo.FullName,
			CreateAt:  repo.CreateAt,
			UpdatedAt: repo.UpdatedAt,
		})
	}
	return response
}
