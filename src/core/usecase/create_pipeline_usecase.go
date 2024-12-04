package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ICreatePipelineUsecase interface {
	CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error)
}
type CreatePipelineUsecase struct {
	thirdPartyToolPort port.IThirdPartyToolPort
}

func (c CreatePipelineUsecase) CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error) {
	//add projectID to pipeline

	newPipeline, err := c.thirdPartyToolPort.CreateNewPipeline(ctx, "demo-ci-cd", pipeline)
	if err != nil {
		log.Error(ctx, "Error when create new pipeline", err)
	}
	actions := pipeline.Actions
	for _, action := range actions {
		_, err := c.thirdPartyToolPort.CreateNewAction(ctx, "test", newPipeline.ID, action)
		if err != nil {
			log.Error(ctx, "Error when create new action", err)
		}
	}
	return newPipeline, nil
}

func NewCreatePipelineUsecase(thirdPartyToolPort port.IThirdPartyToolPort) ICreatePipelineUsecase {
	return &CreatePipelineUsecase{
		thirdPartyToolPort: thirdPartyToolPort,
	}
}
