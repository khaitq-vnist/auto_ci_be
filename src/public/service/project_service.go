package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IProjectService interface {
	GetProjectList(ctx context.Context, userID int64) ([]*entity.ProjectEntity, error)
	CreateProject(ctx context.Context, userID, integrationID, repoID int64) (*entity.ProjectEntity, error)
}
type ProjectService struct {
	getProjectUseCase    usecase.IGetProjectUseCase
	createProjectUseCase usecase.ICreateProjectUseCase
}

func (p ProjectService) CreateProject(ctx context.Context, userID, integrationID, repoID int64) (*entity.ProjectEntity, error) {
	return p.createProjectUseCase.CreateProject(&ctx, userID, integrationID, repoID)
}

func (p ProjectService) GetProjectList(ctx context.Context, userID int64) ([]*entity.ProjectEntity, error) {
	return p.getProjectUseCase.GetProjectList(ctx, userID)
}

func NewProjectService(getProjectUseCase usecase.IGetProjectUseCase, createProjectUseCase usecase.ICreateProjectUseCase) IProjectService {
	return &ProjectService{
		getProjectUseCase:    getProjectUseCase,
		createProjectUseCase: createProjectUseCase,
	}
}
