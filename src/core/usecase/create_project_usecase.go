package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/request"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ICreateProjectUseCase interface {
	CreateProject(ctx context.Context, userId, integrationId, repoId int64) (*entity.ProjectEntity, error)
}
type CreateProjectUseCase struct {
	getRepositoryUseCase  IGetRepositoryUseCase
	projectPort           port.IProjectPort
	getIntegrationUseCase IGetIntegrationUseCase
	thirdPartyToolPort    port.IThirdPartyToolPort
}

func (c CreateProjectUseCase) CreateProject(ctx context.Context, userId, integrationId, repoId int64) (*entity.ProjectEntity, error) {
	integration, err := c.getIntegrationUseCase.GetIntegrationByIdAndUserId(ctx, integrationId, userId)
	if err != nil {
		log.Error(ctx, "get integration error: %v", err)
		return nil, err
	}
	repo, err := c.getRepositoryUseCase.GetRepositoryInfo(ctx, integration.ID, repoId, userId)
	if err != nil {
		log.Error(ctx, "get repository error: %v", err)
		return nil, err
	}
	thirdPartyCreateProjectDto := &request.ThirdPartyCreateProjectRequest{
		DisplayName:       repo.Name,
		Name:              repo.Name,
		Integration:       request.ThirdPartyCreateProjectIntegration{HashID: integration.ThirdPartyHashId},
		ExternalProjectID: repo.FullName,
	}
	thirdPartyCreateProjectRsp, err := c.thirdPartyToolPort.CreateProject(ctx, thirdPartyCreateProjectDto)
	if err != nil {
		log.Error(ctx, "create project error: %v", err)
		return nil, err
	}
	project := &entity.ProjectEntity{
		Name:                repo.Name,
		FullName:            repo.FullName,
		Private:             repo.Private,
		OwnerID:             uint(userId),
		Description:         &repo.Description,
		Language:            &repo.Language,
		ProviderRepoID:      repo.ID,
		HtmlUrl:             repo.HtmlUrl,
		ThirdPartyProjectID: thirdPartyCreateProjectRsp.Name,
		IntegrationId:       integrationId,
	}
	return c.projectPort.SaveProject(ctx, project)
}

func NewCreateProjectUseCase(getRepositoryUseCase IGetRepositoryUseCase, projectPort port.IProjectPort, getIntegrationUseCase IGetIntegrationUseCase, thirdPartyToolPort port.IThirdPartyToolPort) ICreateProjectUseCase {
	return &CreateProjectUseCase{
		getRepositoryUseCase:  getRepositoryUseCase,
		projectPort:           projectPort,
		getIntegrationUseCase: getIntegrationUseCase,
		thirdPartyToolPort:    thirdPartyToolPort,
	}
}
