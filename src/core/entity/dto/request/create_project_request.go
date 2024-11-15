package request

type CreateProjectRequest struct {
	ProviderRepoID int64
	IntegrationID  int64
	OwnerID        int64
	Name           string
	FullName       string
	Private        bool
	Description    string
	Language       string
	ProviderCode   string
}
