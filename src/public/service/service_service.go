package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IServiceService interface {
	GetAllService(ctx context.Context) ([]*entity.ServiceEntity, error)
}
type ServiceService struct {
	getServiceUseCase usecase.IGetServiceUseCase
}

func (s ServiceService) GetAllService(ctx context.Context) ([]*entity.ServiceEntity, error) {
	return s.getServiceUseCase.GetAllService(ctx)
}

func NewServiceService(getServiceUseCase usecase.IGetServiceUseCase) IServiceService {
	return &ServiceService{
		getServiceUseCase: getServiceUseCase,
	}
}
