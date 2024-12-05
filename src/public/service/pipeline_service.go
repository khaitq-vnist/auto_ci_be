package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IPipelineService interface {
	CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error)
	GetListPipelineByProjectID(ctx context.Context, projectID int64) ([]*entity.PipelineEntity, error)
	GetListExecutions(ctx context.Context, projectID, pipelineID int64) (*response.ThirdPartyListExecutionResponse, error)
	GetExecutionDetailByID(ctx context.Context, projectID, pipelineID, executionID int64) (*response.ExecutionResponse, error)
	RunExecution(ctx context.Context, projectID, pipelineID int64) (*response.ExecutionResponse, error)
}
type PipelineService struct {
	createPipelineUsecase usecase.ICreatePipelineUsecase
	getPipelineUseCase    usecase.IGetPipelineUseCase
	getExecutionUsecase   usecase.IGetExecutionUsecase
	runExecutionUsecase   usecase.IRunExecutionUsecase
}

func (p PipelineService) RunExecution(ctx context.Context, projectID, pipelineID int64) (*response.ExecutionResponse, error) {
	return p.runExecutionUsecase.RunExecution(ctx, projectID, pipelineID)
}

func (p PipelineService) GetExecutionDetailByID(ctx context.Context, projectID, pipelineID, executionID int64) (*response.ExecutionResponse, error) {
	return p.getExecutionUsecase.GetExecutionDetailByID(ctx, projectID, pipelineID, executionID)
}

func (p PipelineService) GetListExecutions(ctx context.Context, projectID, pipelineID int64) (*response.ThirdPartyListExecutionResponse, error) {
	return p.getExecutionUsecase.GetListExecutions(ctx, projectID, pipelineID)
}

func (p PipelineService) GetListPipelineByProjectID(ctx context.Context, projectID int64) ([]*entity.PipelineEntity, error) {
	return p.getPipelineUseCase.GetListPipelineOfProject(ctx, projectID)
}

func (p PipelineService) CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error) {
	return p.createPipelineUsecase.CreateNewPipeline(ctx, projectID, pipeline)
}

func NewPipelineService(createPipelineUsecase usecase.ICreatePipelineUsecase, getPipelineUseCase usecase.IGetPipelineUseCase, getExecutionUsecase usecase.IGetExecutionUsecase, runExecutionUsecase usecase.IRunExecutionUsecase) IPipelineService {
	return &PipelineService{
		createPipelineUsecase: createPipelineUsecase,
		getPipelineUseCase:    getPipelineUseCase,
		getExecutionUsecase:   getExecutionUsecase,
		runExecutionUsecase:   runExecutionUsecase,
	}
}
