package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetServiceUseCase interface {
	GetAllService(ctx context.Context) ([]*entity.ServiceEntity, error)
}
type GetServiceUseCase struct {
	servicePort port.IServicePort
}

func (g GetServiceUseCase) GetAllService(ctx context.Context) ([]*entity.ServiceEntity, error) {
	return g.servicePort.GetAllService(ctx)
}

func NewGetServiceUseCase(servicePort port.IServicePort) IGetServiceUseCase {
	return &GetServiceUseCase{
		servicePort: servicePort,
	}
}
