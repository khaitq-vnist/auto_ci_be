package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetIntegrationUseCase interface {
	GetListIntegrationByUserId(ctx context.Context, userId int64) ([]*entity.IntegrationEntity, error)
	GetIntegrationByIdAndUserId(ctx context.Context, integrationId, userId int64) (*entity.IntegrationEntity, error)
}

type GetIntegrationUseCase struct {
	integrationPort port.IIntegrationPort
}

func (g GetIntegrationUseCase) GetIntegrationByIdAndUserId(ctx context.Context, integrationId, userId int64) (*entity.IntegrationEntity, error) {
	return g.integrationPort.GetIntegrationByIdAndUserId(ctx, integrationId, userId)
}

func (g GetIntegrationUseCase) GetListIntegrationByUserId(ctx context.Context, userId int64) ([]*entity.IntegrationEntity, error) {
	return g.integrationPort.GetIntegrationByUserId(ctx, userId)
}

func NewGetIntegrationUseCase(integrationPort port.IIntegrationPort) IGetIntegrationUseCase {
	return &GetIntegrationUseCase{
		integrationPort: integrationPort,
	}
}
