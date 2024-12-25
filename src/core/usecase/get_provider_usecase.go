package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetProviderUseCase interface {
	GetProviderByCode(ctx context.Context, code string) (*entity.ProviderEntity, error)
}
type GetProviderUseCase struct {
	providerPort port.IProviderPort
}

func (g *GetProviderUseCase) GetProviderByCode(ctx context.Context, code string) (*entity.ProviderEntity, error) {
	return g.providerPort.GetProviderByCode(ctx, code)
}

func NewGetProviderUseCase(providerPort port.IProviderPort) IGetProviderUseCase {
	return &GetProviderUseCase{
		providerPort: providerPort,
	}
}
