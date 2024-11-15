package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type ProjectResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"full_name"`
	UpdatedAt int64  `json:"updated_at"`
}

func ToProjectResponse(project *entity.ProjectEntity) *ProjectResponse {
	return &ProjectResponse{
		ID:        project.ID,
		Name:      project.Name,
		FullName:  project.FullName,
		UpdatedAt: project.UpdatedAt,
	}
}
func ToListProjectResponse(projects []*entity.ProjectEntity) []*ProjectResponse {
	var response []*ProjectResponse
	for _, project := range projects {
		response = append(response, ToProjectResponse(project))
	}
	return response
}
