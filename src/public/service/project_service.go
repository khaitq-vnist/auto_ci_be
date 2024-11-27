package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IProjectService interface {
	GetProjectList(ctx context.Context, userID int64) ([]*entity.ProjectEntity, error)
	CreateProject(ctx context.Context, userID, integrationID, repoID int64) (*entity.ProjectEntity, error)
	AnalyzeProject(ctx context.Context, userID, projectID int64) (*dto.AnalyzeSourceCodeDTO, error)
	GetListBranches(ctx context.Context, userID, projectID int64) ([]*response.ThirdPartyBranchResponse, error)
}
type ProjectService struct {
	getProjectUseCase        usecase.IGetProjectUseCase
	createProjectUseCase     usecase.ICreateProjectUseCase
	analyzeSourceCodeUsecase usecase.IAnalyzeSourceCodeUsecase
	getListBranchUseCase     usecase.IGetBranchUseCase
}

func (p ProjectService) GetListBranches(ctx context.Context, userID, projectID int64) ([]*response.ThirdPartyBranchResponse, error) {
	return p.getListBranchUseCase.GetListBranches(ctx, userID, projectID)
}

func (p ProjectService) AnalyzeProject(ctx context.Context, userID, projectID int64) (*dto.AnalyzeSourceCodeDTO, error) {
	return p.analyzeSourceCodeUsecase.Analyze(ctx, userID, projectID)
}

func (p ProjectService) CreateProject(ctx context.Context, userID, integrationID, repoID int64) (*entity.ProjectEntity, error) {
	return p.createProjectUseCase.CreateProject(&ctx, userID, integrationID, repoID)
}

func (p ProjectService) GetProjectList(ctx context.Context, userID int64) ([]*entity.ProjectEntity, error) {
	return p.getProjectUseCase.GetProjectList(ctx, userID)
}

func NewProjectService(getProjectUseCase usecase.IGetProjectUseCase, createProjectUseCase usecase.ICreateProjectUseCase, analyzeSourceCodeUsecase usecase.IAnalyzeSourceCodeUsecase, getListBranchUseCase usecase.IGetBranchUseCase) IProjectService {
	return &ProjectService{
		getProjectUseCase:        getProjectUseCase,
		createProjectUseCase:     createProjectUseCase,
		analyzeSourceCodeUsecase: analyzeSourceCodeUsecase,
		getListBranchUseCase:     getListBranchUseCase,
	}
}
