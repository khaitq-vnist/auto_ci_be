package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetProjectUseCase interface {
	GetProjectList(c context.Context, userId int64) ([]*entity.ProjectEntity, error)
	GetProjectById(c context.Context, projectId int64) (*entity.ProjectEntity, error)
}
type GetProjectUseCase struct {
	projectPort port.IProjectPort
}

func (g GetProjectUseCase) GetProjectById(c context.Context, projectId int64) (*entity.ProjectEntity, error) {
	return g.projectPort.GetProjectById(c, projectId)
}

func (g GetProjectUseCase) GetProjectList(c context.Context, userId int64) ([]*entity.ProjectEntity, error) {
	return g.projectPort.GetListProjectByUserId(c, userId)
}

func NewGetProjectUseCase(projectPort port.IProjectPort) IGetProjectUseCase {
	return &GetProjectUseCase{
		projectPort: projectPort,
	}
}
