package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/request"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
)

type IThirdPartyToolPort interface {
	CreateNewPipeline(ctx context.Context, project string, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error)
	CreateNewAction(ctx context.Context, project string, pipelineID int64, action *entity.ActionEntity) (*entity.ActionEntity, error)
	GetListPipeline(ctx context.Context, project string) ([]*entity.PipelineEntity, error)
	GetListExecutions(ctx context.Context, project string, pipelineID int64) (*response.ThirdPartyListExecutionResponse, error)
	GetExecutionDetail(ctx context.Context, project string, pipelineID, executionID int64) (*response.ExecutionResponse, error)
	RunExecution(ctx context.Context, project string, pipelineID int64) (*response.ExecutionResponse, error)
	DeletePipelineById(ctx context.Context, project string, pipelineID int64) error
	GetDetailLog(ctx context.Context, project string, pipelineID, executionID, actionId int64) (*response.DetailActionLog, error)
	CreateIntegration(ctx context.Context, integration *entity.IntegrationEntity) (*response.ThirdPartyCreateIntegrationResponse, error)
	CreateProject(ctx context.Context, projectDto *request.ThirdPartyCreateProjectRequest) (*response.ThirdPartyCreateProjectResponse, error)
}
