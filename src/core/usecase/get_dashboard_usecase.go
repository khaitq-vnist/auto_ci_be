package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetDashboardUsecase interface {
	GetDashboard(ctx context.Context, userID int64) (*response.DashboardResponse, error)
}
type GetDashboardUsecase struct {
	projectPort     port.IProjectPort
	integrationPort port.IIntegrationPort
}

func (g GetDashboardUsecase) GetDashboard(ctx context.Context, userID int64) (*response.DashboardResponse, error) {
	totalProjects, err := g.projectPort.CountAllProjectByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}
	totalIntegrations, err := g.integrationPort.CountAllIntegrationByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &response.DashboardResponse{
		TotalProjects:     totalProjects,
		TotalIntegrations: totalIntegrations,
	}, nil
}

func NewGetDashboardUsecase(projectPort port.IProjectPort, integrationPort port.IIntegrationPort) IGetDashboardUsecase {
	return &GetDashboardUsecase{
		projectPort:     projectPort,
		integrationPort: integrationPort,
	}
}
