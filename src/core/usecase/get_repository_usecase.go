package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetRepositoryUseCase interface {
	GetReposByIntegrationId(ctx context.Context, integrationId, userId int64) ([]*response.ThirdPartyProviderReposResponse, error)
	GetRepositoryInfo(ctx context.Context, integrationId, repoId int64) (*response.ThirdPartyProviderReposResponse, error)
}

type GetRepositoryUseCase struct {
	thirdPartyPort        port.IThirdPartyProviderPort
	getIntegrationUseCase IGetIntegrationUseCase
	encryptUseCase        IEncryptUseCase
}

func (g GetRepositoryUseCase) GetRepositoryInfo(ctx context.Context, integrationId, repoId int64) (*response.ThirdPartyProviderReposResponse, error) {
	integration, err := g.getIntegrationUseCase.GetIntegrationByIdAndUserId(ctx, integrationId, 1)
	if err != nil {
		return nil, err
	}
	if integration.ProviderName == "GitHub" {
		integration.ProviderName = "GITHUB"
	}
	decryptToken, err := g.encryptUseCase.DecryptToken(&ctx, integration.AccessToken)
	if err != nil {
		return nil, err
	}
	return g.thirdPartyPort.GetRepositoryInfo(&ctx, integration.ProviderName, decryptToken, repoId)
}

func (g GetRepositoryUseCase) GetReposByIntegrationId(ctx context.Context, integrationId, userId int64) ([]*response.ThirdPartyProviderReposResponse, error) {
	integration, err := g.getIntegrationUseCase.GetIntegrationByIdAndUserId(ctx, integrationId, 1)
	if err != nil {
		return nil, err
	}
	if integration.ProviderName == "GitHub" {
		integration.ProviderName = "GITHUB"
	}
	decryptToken, err := g.encryptUseCase.DecryptToken(&ctx, integration.AccessToken)
	if err != nil {
		return nil, err
	}

	return g.thirdPartyPort.GetListRepositoriesByUser(&ctx, integration.ProviderName, decryptToken, integration.ProviderUsername)
}

func NewGetRepositoryUseCase(thirdPartyPort port.IThirdPartyProviderPort, getIntegrationUseCase IGetIntegrationUseCase, encryptUseCase IEncryptUseCase) IGetRepositoryUseCase {
	return &GetRepositoryUseCase{
		thirdPartyPort:        thirdPartyPort,
		getIntegrationUseCase: getIntegrationUseCase,
		encryptUseCase:        encryptUseCase,
	}
}
