package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetPipelineUseCase interface {
	GetListPipelineOfProject(ctx context.Context, projectId int64) ([]*entity.PipelineEntity, error)
}

type GetPipelineUseCase struct {
	getProjectUseCase  IGetProjectUseCase
	thirdPartyToolPort port.IThirdPartyToolPort
}

func (g GetPipelineUseCase) GetListPipelineOfProject(ctx context.Context, projectId int64) ([]*entity.PipelineEntity, error) {
	project, err := g.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		log.Error(ctx, "Error when get project by id", err)
		return nil, err
	}

	pipelines, err := g.thirdPartyToolPort.GetListPipeline(ctx, project.ThirdPartyProjectID)
	if err != nil {
		log.Error(ctx, "Error when get list pipeline", err)
		return nil, err
	}
	return pipelines, nil
}

func NewGetPipelineUseCase(getProjectUseCase IGetProjectUseCase, thirdPartyToolPort port.IThirdPartyToolPort) IGetPipelineUseCase {
	return &GetPipelineUseCase{
		getProjectUseCase:  getProjectUseCase,
		thirdPartyToolPort: thirdPartyToolPort,
	}
}
