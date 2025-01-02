package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/constant"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ICreateIntegrationUseCase interface {
	CreateIntegration(ctx context.Context, integrationEntity *entity.IntegrationEntity) (*entity.IntegrationEntity, error)
}
type CreateIntegrationUseCase struct {
	encryptUseCase               IEncryptUseCase
	getThirdPartyProviderUseCase IGetThirdPartyProviderUseCase
	getProviderUseCase           IGetProviderUseCase
	integrationPort              port.IIntegrationPort
	thirdPartyToolPort           port.IThirdPartyToolPort
}

func (c *CreateIntegrationUseCase) CreateIntegration(ctx context.Context, integrationEntity *entity.IntegrationEntity) (*entity.IntegrationEntity, error) {
	provider, err := c.getProviderUseCase.GetProviderByCode(ctx, integrationEntity.ProviderCode)
	if err != nil {
		log.Error(ctx, "Error when get provider by code", err)
		return nil, err
	}
	userInfo, err := c.getThirdPartyProviderUseCase.GetUserInfo(ctx, provider.Code, integrationEntity.AccessToken)
	if err != nil {
		log.Error(ctx, "Error when get user info from third party provider", err)
		return nil, err
	}
	if userInfo == nil {
		log.Error(ctx, "User info not found", nil)
		return nil, err
	}
	var providerType string
	if provider.Code == constant.GITHUB_PROVIDER {
		providerType = constant.GITHUB_INTEGRATION
	}
	integrationEntity.Type = providerType
	integrationEntity.Scope = constant.WORKSPACE_SCOPE
	thirdPartyIntegration, err := c.thirdPartyToolPort.CreateIntegration(ctx, integrationEntity)

	if err != nil {
		log.Error(ctx, "Error when create integration", err)
		return nil, err
	}
	integrationEntity.ThirdPartyHashId = thirdPartyIntegration.HashID

	encryptToken, err := c.encryptUseCase.EncryptToken(ctx, integrationEntity.AccessToken)
	if err != nil {
		log.Error(ctx, "Error when encrypt token", err)
		return nil, err
	}

	integrationEntity.AccessToken = encryptToken
	integrationEntity.ProviderName = provider.Code
	integrationEntity.ProviderId = provider.ID
	integrationEntity.ProviderUsername = userInfo.Username
	integrationEntity.ProviderCode = provider.Code

	integrationEntity, err = c.integrationPort.SaveIntegration(ctx, integrationEntity)
	if err != nil {
		log.Error(ctx, "Error when save integration", err)
		return nil, err
	}
	integrationEntity.AccessToken = ""

	return integrationEntity, nil

}

func NewCreateIntegrationUseCase(encryptUseCase IEncryptUseCase, getThirdPartyProviderUseCase IGetThirdPartyProviderUseCase, getProviderUseCase IGetProviderUseCase, integrationPort port.IIntegrationPort, thirdPartyToolPort port.IThirdPartyToolPort) ICreateIntegrationUseCase {
	return &CreateIntegrationUseCase{
		encryptUseCase:               encryptUseCase,
		getThirdPartyProviderUseCase: getThirdPartyProviderUseCase,
		getProviderUseCase:           getProviderUseCase,
		integrationPort:              integrationPort,
		thirdPartyToolPort:           thirdPartyToolPort,
	}
}
