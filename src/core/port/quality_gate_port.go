package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
)

type IQualityGatePort interface {
	CreateNewProject(ctx context.Context, name, project string) (*response.SonarCreateProjectResponse, error)
	GetAccessToken(ctx context.Context, name, projectKey string) (string, error)
	GetCoverage(ctx context.Context, projectKey string) (float64, error)
}
