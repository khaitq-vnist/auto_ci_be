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

	newPipeline, err := c.thirdPartyToolPort.CreateNewPipeline(ctx, project.ThirdPartyProjectID, pipeline)
	if err != nil {
		log.Error(ctx, "Error when create new pipeline", err)
		return nil, err
	}
	actions := pipeline.Actions
	for _, action := range actions {
		if len(action.ExecuteCommands) > 0 && action.ExecuteCommands[0] == "mvn clean verify" {
			action.ExecuteCommands = append(action.ExecuteCommands,
				"mvn sonar:sonar "+
					"-Dsonar.projectKey="+project.SonarKey+" "+
					"-Dsonar.projectName="+project.SonarProjectName+" "+
					"-Dsonar.host.url=https://sonar.auto-ci.site "+
					"-Dsonar.token="+project.SonarToken,
			)

		}
		_, err := c.thirdPartyToolPort.CreateNewAction(ctx, project.ThirdPartyProjectID, newPipeline.ID, action)
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
