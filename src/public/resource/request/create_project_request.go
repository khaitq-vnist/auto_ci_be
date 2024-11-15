package request

type CreateProjectRequest struct {
	IntegrationId int64 `json:"integration_id"`
	RepoId        int64 `json:"repo_id"`
}
