package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IPipelineService interface {
	CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error)
}
type PipelineService struct {
	createPipelineUsecase usecase.ICreatePipelineUsecase
}

func (p PipelineService) CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error) {
	return p.createPipelineUsecase.CreateNewPipeline(ctx, projectID, pipeline)
}

func NewPipelineService(createPipelineUsecase usecase.ICreatePipelineUsecase) IPipelineService {
	return &PipelineService{
		createPipelineUsecase: createPipelineUsecase,
	}
}
