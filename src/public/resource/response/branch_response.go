package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"

type BranchResponse struct {
	Name   string         `json:"name"`
	Commit CommitResponse `json:"commit"`
}
type CommitResponse struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

func ToListBranchResponse(branches []*response.ThirdPartyBranchResponse) []*BranchResponse {
	var result []*BranchResponse
	for _, branch := range branches {
		result = append(result, &BranchResponse{
			Name: branch.Name,
			Commit: CommitResponse{
				Sha: branch.Commit.Sha,
				URL: branch.Commit.URL,
			},
		})
	}
	return result
}
