package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetExecutionUsecase interface {
	GetListExecutions(ctx context.Context, projectId, pipelineID int64) (*response.ThirdPartyListExecutionResponse, error)
	GetExecutionDetailByID(ctx context.Context, projectId, pipelineID, executionID int64) (*response.ExecutionResponse, error)
}
type GetExecutionUsecase struct {
	getProjectUseCase  IGetProjectUseCase
	thirdPartyToolPort port.IThirdPartyToolPort
}

func (g GetExecutionUsecase) GetExecutionDetailByID(ctx context.Context, projectId, pipelineID, executionID int64) (*response.ExecutionResponse, error) {
	project, err := g.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		return nil, err
	}
	project.Name = "demo-ci-cd"
	return g.thirdPartyToolPort.GetExecutionDetail(ctx, project.Name, pipelineID, executionID)
}

func (g GetExecutionUsecase) GetListExecutions(ctx context.Context, projectId, pipelineID int64) (*response.ThirdPartyListExecutionResponse, error) {
	project, err := g.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		return nil, err
	}
	project.Name = "demo-ci-cd"
	return g.thirdPartyToolPort.GetListExecutions(ctx, project.Name, pipelineID)
}

func NewGetExecutionUsecase(getProjectUseCase IGetProjectUseCase, thirdPartyToolPort port.IThirdPartyToolPort) IGetExecutionUsecase {
	return &GetExecutionUsecase{
		getProjectUseCase:  getProjectUseCase,
		thirdPartyToolPort: thirdPartyToolPort,
	}
}
