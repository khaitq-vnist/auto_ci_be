package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetBranchUseCase interface {
	GetListBranches(ctx context.Context, userId, projectId int64) ([]*response.ThirdPartyBranchResponse, error)
}

type GetBranchUseCase struct {
	thirdPartyProviderPort port.IThirdPartyProviderPort
	getIntegrationUseCase  IGetIntegrationUseCase
	getProjectUseCase      IGetProjectUseCase
	encryptUseCase         IEncryptUseCase
}

func (g GetBranchUseCase) GetListBranches(ctx context.Context, userId, projectId int64) ([]*response.ThirdPartyBranchResponse, error) {
	project, err := g.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		log.Error(ctx, "Get project failed", err)
		return nil, err
	}

	integration, err := g.getIntegrationUseCase.GetIntegrationByIdAndUserId(ctx, userId, project.IntegrationId)
	if err != nil {
		log.Error(ctx, "Get integration failed", err)
		return nil, err
	}

	decryptToken, err := g.encryptUseCase.DecryptToken(ctx, integration.AccessToken)
	if err != nil {
		log.Error(ctx, "Decrypt token failed", err)
		return nil, err
	}
	branches, err := g.thirdPartyProviderPort.GetListBranches(&ctx, integration.ProviderName, integration.ProviderUsername, decryptToken, project.Name)
	if err != nil {
		log.Error(ctx, "Get list branch failed", err)
		return nil, err
	}
	return branches, nil
}

func NewGetBranchUseCase(thirdPartyProviderPort port.IThirdPartyProviderPort, getIntegrationUseCase IGetIntegrationUseCase, getProjectUseCase IGetProjectUseCase, encryptUseCase IEncryptUseCase) IGetBranchUseCase {
	return &GetBranchUseCase{
		thirdPartyProviderPort: thirdPartyProviderPort,
		getIntegrationUseCase:  getIntegrationUseCase,
		getProjectUseCase:      getProjectUseCase,
		encryptUseCase:         encryptUseCase,
	}
}
