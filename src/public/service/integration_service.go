package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/request"
)

type IIntegrationService interface {
	CreateIntegration(c context.Context, userID int64, request *request.CreateIntegrationRequest) error
	GetIntegrationByUserId(c context.Context, userId int64) ([]*entity.IntegrationEntity, error)
}

type IntegrationService struct {
	createIntegrationUseCase usecase.ICreateIntegrationUseCase
	getIntegrationUseCase    usecase.IGetIntegrationUseCase
}

func (i *IntegrationService) GetIntegrationByUserId(c context.Context, userId int64) ([]*entity.IntegrationEntity, error) {
	return i.getIntegrationUseCase.GetListIntegrationByUserId(c, userId)
}

func (i *IntegrationService) CreateIntegration(c context.Context, userID int64, request *request.CreateIntegrationRequest) error {
	integrationEntity := request.ToEntity()
	integrationEntity.UserId = userID
	_, err := i.createIntegrationUseCase.CreateIntegration(c, integrationEntity)
	return err
}

func NewIntegrationService(createIntegrationUseCase usecase.ICreateIntegrationUseCase, getIntegrationUseCase usecase.IGetIntegrationUseCase) IIntegrationService {
	return &IntegrationService{
		createIntegrationUseCase: createIntegrationUseCase,
		getIntegrationUseCase:    getIntegrationUseCase,
	}
}
