package postgres

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"gorm.io/gorm"
)

type ExecutionHistoryRepositoryAdapter struct {
	BaseRepository
}

func (e ExecutionHistoryRepositoryAdapter) GetExecutionHistoryByProjectIDAndPipelineIDExecutionIDs(ctx context.Context, projectID, pipelineID int64, executionIDs []int64) ([]*entity.ExecutionHistoryEntity, error) {
	var hisModels []*model.ExecutionHistoryModel
	if err := e.db.WithContext(ctx).Model(&model.ExecutionHistoryModel{}).Where("project_id = ? AND pipeline_id = ? AND third_party_id IN ?", projectID, pipelineID, executionIDs).Find(&hisModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListExecutionHistoryEntity(hisModels), nil
}

func (e ExecutionHistoryRepositoryAdapter) GetExecutionHistoryByProjectIDAndPipelineIDExecutionID(ctx context.Context, projectID, pipelineID, executionID int64) (*entity.ExecutionHistoryEntity, error) {
	var hisModel model.ExecutionHistoryModel
	if err := e.db.WithContext(ctx).Model(&model.ExecutionHistoryModel{}).Where("project_id = ? AND pipeline_id = ? AND third_party_id = ?", projectID, pipelineID, executionID).First(&hisModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToExecutionHistoryEntity(&hisModel), nil
}

func (e ExecutionHistoryRepositoryAdapter) CreateExecutionHistory(ctx context.Context, his *entity.ExecutionHistoryEntity) (*entity.ExecutionHistoryEntity, error) {
	hisModel := mapper.ToExecutionHistoryModel(his)
	if err := e.db.WithContext(ctx).Model(&model.ExecutionHistoryModel{}).Create(&hisModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToExecutionHistoryEntity(hisModel), nil
}

func NewExecutionHistoryRepositoryAdapter(db *gorm.DB) port.IExecutionHistoryPort {
	return &ExecutionHistoryRepositoryAdapter{
		BaseRepository: BaseRepository{db: db},
	}
}
