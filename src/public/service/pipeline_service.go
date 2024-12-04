package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IPipelineService interface {
	CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error)
	GetListPipelineByProjectID(ctx context.Context, projectID int64) ([]*entity.PipelineEntity, error)
}
type PipelineService struct {
	createPipelineUsecase usecase.ICreatePipelineUsecase
	getPipelineUseCase    usecase.IGetPipelineUseCase
}

func (p PipelineService) GetListPipelineByProjectID(ctx context.Context, projectID int64) ([]*entity.PipelineEntity, error) {
	return p.getPipelineUseCase.GetListPipelineOfProject(ctx, projectID)
}

func (p PipelineService) CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error) {
	return p.createPipelineUsecase.CreateNewPipeline(ctx, projectID, pipeline)
}

func NewPipelineService(createPipelineUsecase usecase.ICreatePipelineUsecase, getPipelineUseCase usecase.IGetPipelineUseCase) IPipelineService {
	return &PipelineService{
		createPipelineUsecase: createPipelineUsecase,
		getPipelineUseCase:    getPipelineUseCase,
	}
}
