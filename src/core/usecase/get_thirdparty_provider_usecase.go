package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetThirdPartyProviderUseCase interface {
	GetUserInfo(ctx *context.Context, provider string, token string) (*dto.ThirdPartyProviderUserInfoResponse, error)
}
type GetThirdPartyProviderUseCase struct {
	thirdPartyProviderPort port.IThirdPartyProviderPort
}

func (g GetThirdPartyProviderUseCase) GetUserInfo(ctx *context.Context, provider string, token string) (*dto.ThirdPartyProviderUserInfoResponse, error) {
	return g.thirdPartyProviderPort.GetUserInfo(ctx, provider, token)
}

func NewGetThirdPartyProviderUseCase(thirdPartyProviderPort port.IThirdPartyProviderPort) IGetThirdPartyProviderUseCase {
	return &GetThirdPartyProviderUseCase{
		thirdPartyProviderPort: thirdPartyProviderPort,
	}
}
