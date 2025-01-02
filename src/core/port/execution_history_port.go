package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IExecutionHistoryPort interface {
	CreateExecutionHistory(ctx context.Context, his *entity.ExecutionHistoryEntity) (*entity.ExecutionHistoryEntity, error)
	GetExecutionHistoryByProjectIDAndPipelineIDExecutionID(ctx context.Context, projectID, pipelineID, executionID int64) (*entity.ExecutionHistoryEntity, error)
	GetExecutionHistoryByProjectIDAndPipelineIDExecutionIDs(ctx context.Context, projectID, pipelineID int64, executionIDs []int64) ([]*entity.ExecutionHistoryEntity, error)
}
