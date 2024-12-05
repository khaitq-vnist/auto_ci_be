package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IRunExecutionUsecase interface {
	RunExecution(ctx context.Context, projectId, pipelineID int64) (*response.ExecutionResponse, error)
}
type RunExecutionUsecase struct {
	getProjectUseCase  IGetProjectUseCase
	thirdPartyToolPort port.IThirdPartyToolPort
}

func (r RunExecutionUsecase) RunExecution(ctx context.Context, projectId, pipelineID int64) (*response.ExecutionResponse, error) {
	project, err := r.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		return nil, err
	}
	project.Name = "demo-ci-cd"
	return r.thirdPartyToolPort.RunExecution(ctx, project.Name, pipelineID)
}

func NewRunExecutionUsecase(getProjectUseCase IGetProjectUseCase, thirdPartyToolPort port.IThirdPartyToolPort) IRunExecutionUsecase {
	return &RunExecutionUsecase{
		getProjectUseCase:  getProjectUseCase,
		thirdPartyToolPort: thirdPartyToolPort,
	}
}
