package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IDashboardService interface {
	GetDashboard(ctx context.Context, userID int64) (*response.DashboardResponse, error)
}
type DashboardService struct {
	getDashboardUsecase usecase.IGetDashboardUsecase
}

func (d DashboardService) GetDashboard(ctx context.Context, userID int64) (*response.DashboardResponse, error) {
	return d.getDashboardUsecase.GetDashboard(ctx, userID)
}

func NewDashboardService(getDashboardUsecase usecase.IGetDashboardUsecase) IDashboardService {
	return &DashboardService{
		getDashboardUsecase: getDashboardUsecase,
	}
}
