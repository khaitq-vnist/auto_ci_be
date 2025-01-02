package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"strings"
)

type IGetExecutionUsecase interface {
	GetListExecutions(ctx context.Context, projectId, pipelineID int64) (*response.ThirdPartyListExecutionResponse, error)
	GetExecutionDetailByID(ctx context.Context, projectId, pipelineID, executionID int64) (*response.ExecutionResponse, error)
}
type GetExecutionUsecase struct {
	getProjectUseCase    IGetProjectUseCase
	thirdPartyToolPort   port.IThirdPartyToolPort
	executionHistoryPort port.IExecutionHistoryPort
}

func (g GetExecutionUsecase) GetExecutionDetailByID(ctx context.Context, projectId, pipelineID, executionID int64) (*response.ExecutionResponse, error) {
	project, err := g.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		return nil, err
	}

	result, err := g.thirdPartyToolPort.GetExecutionDetail(ctx, project.ThirdPartyProjectID, pipelineID, executionID)
	if err != nil {
		log.Error(ctx, "GetExecutionDetail error", err)
		return nil, err
	}
	exhist, err := g.executionHistoryPort.GetExecutionHistoryByProjectIDAndPipelineIDExecutionID(ctx, projectId, pipelineID, executionID)
	if err != nil && err.Error() != "record not found" {
		log.Error(ctx, "GetExecutionHistoryByProjectIDAndPipelineIDExecutionID error", err)
		return nil, err
	}
	if exhist != nil {
		logs := strings.Split(exhist.LogsFile, ",")
		result.LogsFile = logs
		result.Coverage = exhist.Coverage
	}
	return result, nil

}

func (g GetExecutionUsecase) GetListExecutions(ctx context.Context, projectId, pipelineID int64) (*response.ThirdPartyListExecutionResponse, error) {
	project, err := g.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		return nil, err
	}

	results, err := g.thirdPartyToolPort.GetListExecutions(ctx, project.ThirdPartyProjectID, pipelineID)
	if err != nil {
		log.Error(ctx, "GetListExecutions error", err)
		return nil, err
	}
	exectionIds := make([]int64, 0)
	for _, execution := range results.Executions {
		exectionIds = append(exectionIds, execution.ID)
	}
	exhist, err := g.executionHistoryPort.GetExecutionHistoryByProjectIDAndPipelineIDExecutionIDs(ctx, projectId, pipelineID, exectionIds)
	if err != nil {
		log.Error(ctx, "GetExecutionHistoryByProjectIDAndPipelineIDExecutionIDs error", err)
		return nil, err
	}
	exHistMap := make(map[int64]*entity.ExecutionHistoryEntity)
	for _, ex := range exhist {
		exHistMap[ex.ThirdPartyID] = ex
	}
	for i, execution := range results.Executions {
		if ex, ok := exHistMap[execution.ID]; ok {
			logs := strings.Split(ex.LogsFile, ",")
			results.Executions[i].LogsFile = logs
			results.Executions[i].Coverage = ex.Coverage
		}
	}
	return results, nil
}

func NewGetExecutionUsecase(getProjectUseCase IGetProjectUseCase, thirdPartyToolPort port.IThirdPartyToolPort, executionHistoryPort port.IExecutionHistoryPort) IGetExecutionUsecase {
	return &GetExecutionUsecase{
		getProjectUseCase:    getProjectUseCase,
		thirdPartyToolPort:   thirdPartyToolPort,
		executionHistoryPort: executionHistoryPort,
	}
}
