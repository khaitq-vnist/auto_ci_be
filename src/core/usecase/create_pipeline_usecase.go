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
	getProjectUseCase     IGetProjectUseCase
	thirdPartyToolPort    port.IThirdPartyToolPort
	deletePipelineUsecase IDeletePipelineUsecase
}

func (c CreatePipelineUsecase) CreateNewPipeline(ctx context.Context, projectID int64, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error) {
	//add projectID to pipeline
	project, err := c.getProjectUseCase.GetProjectById(ctx, projectID)
	if err != nil {
		log.Error(ctx, "Error when get project by id", err)
		return nil, err
	}
	project.Name = "demo-ci-cd"
	newPipeline, err := c.thirdPartyToolPort.CreateNewPipeline(ctx, "demo-ci-cd", pipeline)
	if err != nil {
		log.Error(ctx, "Error when create new pipeline", err)
		return nil, err
	}
	actions := pipeline.Actions
	for _, action := range actions {
		_, err := c.thirdPartyToolPort.CreateNewAction(ctx, "demo-ci-cd", newPipeline.ID, action)
		if err != nil {
			log.Error(ctx, "Error when create new action", err)
			go func() {
				_ = c.deletePipelineUsecase.DeletePipeline(ctx, projectID, newPipeline.ID)
				log.Info(ctx, "Delete pipeline after create action failed")
			}()
			return nil, err
		}
	}
	return newPipeline, nil
}

func NewCreatePipelineUsecase(getProjectUseCase IGetProjectUseCase, thirdPartyToolPort port.IThirdPartyToolPort, deletePipelineUsecase IDeletePipelineUsecase) ICreatePipelineUsecase {
	return &CreatePipelineUsecase{
		getProjectUseCase:     getProjectUseCase,
		thirdPartyToolPort:    thirdPartyToolPort,
		deletePipelineUsecase: deletePipelineUsecase,
	}
}
