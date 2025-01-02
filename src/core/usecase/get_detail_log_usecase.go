package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetDetailLogUseCase interface {
	GetDetailLog(ctx context.Context, projectID, pipelineID, executionID, actionID int64) (*response.DetailActionLog, error)
}
type GetDetailLogUseCase struct {
	thirdPartyToolPort port.IThirdPartyToolPort
	getProjectUseCase  IGetProjectUseCase
}

func (g GetDetailLogUseCase) GetDetailLog(ctx context.Context, projectID, pipelineID, executionID, actionID int64) (*response.DetailActionLog, error) {
	project, err := g.getProjectUseCase.GetProjectById(ctx, projectID)
	if err != nil {
		log.Error(ctx, "Error when get project by id", err)
		return nil, err
	}
	return g.thirdPartyToolPort.GetDetailLog(ctx, project.ThirdPartyProjectID, pipelineID, executionID, actionID)
}

func NewGetDetailLogUseCase(thirdPartyToolPort port.IThirdPartyToolPort, getProjectUseCase IGetProjectUseCase) IGetDetailLogUseCase {
	return &GetDetailLogUseCase{
		thirdPartyToolPort: thirdPartyToolPort,
		getProjectUseCase:  getProjectUseCase,
	}
}
