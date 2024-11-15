package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IRepositoryService interface {
	GetRepositoriesByIntegrationId(c context.Context, inId, userId int64) ([]*response.ThirdPartyProviderReposResponse, error)
}
type RepositoryService struct {
	getRepositoryUseCase usecase.IGetRepositoryUseCase
}

func (r RepositoryService) GetRepositoriesByIntegrationId(c context.Context, itId, userId int64) ([]*response.ThirdPartyProviderReposResponse, error) {
	return r.getRepositoryUseCase.GetReposByIntegrationId(c, itId, userId)
}

func NewRepositoryService(getRepositoryUseCase usecase.IGetRepositoryUseCase) IRepositoryService {
	return &RepositoryService{
		getRepositoryUseCase: getRepositoryUseCase,
	}
}
