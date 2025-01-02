package request

type SonarCreateProjectRequest struct {
	Name    string `json:"name"`
	Project string `json:"project"`
}

func ToSonarCreateProjectRequest(name, project string) *SonarCreateProjectRequest {
	return &SonarCreateProjectRequest{
		Name:    name,
		Project: project,
	}
}
