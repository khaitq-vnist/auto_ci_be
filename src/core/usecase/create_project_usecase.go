package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ICreateProjectUseCase interface {
	CreateProject(ctx *context.Context, userId, integrationId, repoId int64) (*entity.ProjectEntity, error)
}
type CreateProjectUseCase struct {
	getRepositoryUseCase IGetRepositoryUseCase
	projectPort          port.IProjectPort
}

func (c CreateProjectUseCase) CreateProject(ctx *context.Context, userId, integrationId, repoId int64) (*entity.ProjectEntity, error) {
	repo, err := c.getRepositoryUseCase.GetRepositoryInfo(*ctx, integrationId, repoId)
	if err != nil {
		return nil, err
	}
	project := &entity.ProjectEntity{
		Name:           repo.Name,
		FullName:       repo.FullName,
		Private:        repo.Private,
		OwnerID:        uint(userId),
		Description:    &repo.Description,
		Language:       &repo.Language,
		ProviderRepoID: repo.ID,
		HtmlUrl:        repo.HtmlUrl,
	}
	return c.projectPort.SaveProject(*ctx, project)
}

func NewCreateProjectUseCase(getRepositoryUseCase IGetRepositoryUseCase, projectPort port.IProjectPort) ICreateProjectUseCase {
	return &CreateProjectUseCase{
		getRepositoryUseCase: getRepositoryUseCase,
		projectPort:          projectPort,
	}
}
