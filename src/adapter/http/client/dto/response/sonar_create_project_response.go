package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"

type SonarCreateProjectResponse struct {
	Project *SonarCreateProject `json:"project"`
}
type SonarCreateProject struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Qualifier string `json:"qualifier"`
}

func ToSonarCreateProjectResponse(rsp *SonarCreateProjectResponse) *response.SonarCreateProjectResponse {
	if rsp == nil || rsp.Project == nil {
		return nil
	}
	return &response.SonarCreateProjectResponse{
		Key:       rsp.Project.Key,
		Name:      rsp.Project.Name,
		Qualifier: rsp.Project.Qualifier,
	}
}
