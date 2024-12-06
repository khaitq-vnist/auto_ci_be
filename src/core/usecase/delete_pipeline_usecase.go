package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IDeletePipelineUsecase interface {
	DeletePipeline(ctx context.Context, projectID, pipelineID int64) error
}
type DeletePipelineUsecase struct {
	getProjectUseCase  IGetProjectUseCase
	thirdPartyToolPort port.IThirdPartyToolPort
}

func (d DeletePipelineUsecase) DeletePipeline(ctx context.Context, projectID, pipelineID int64) error {
	project, err := d.getProjectUseCase.GetProjectById(ctx, projectID)
	if err != nil {
		log.Error(ctx, "Error when get project by id", err)
		return err
	}
	project.Name = "demo-ci-cd"
	err = d.thirdPartyToolPort.DeletePipelineById(ctx, project.Name, pipelineID)
	if err != nil {
		log.Error(ctx, "Error when delete pipeline", err)
		return err
	}
	return nil
}

func NewDeletePipelineUsecase(getProjectUseCase IGetProjectUseCase, thirdPartyToolPort port.IThirdPartyToolPort) IDeletePipelineUsecase {
	return &DeletePipelineUsecase{
		getProjectUseCase:  getProjectUseCase,
		thirdPartyToolPort: thirdPartyToolPort,
	}
}
