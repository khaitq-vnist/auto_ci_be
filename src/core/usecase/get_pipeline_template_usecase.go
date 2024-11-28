package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetPipelineTemplateUsecase interface {
	GetPipelineTemplate(ctx context.Context, buildTool string) (*response.PipelineTemplateResponse, error)
}

type GetPipelineTemplateUsecase struct {
	pipelineTemplatePort      port.IPipelineTemplatePort
	pipelineStageTemplatePort port.IPipelineStageTemplatePort
	stageTemplatePort         port.IStageTemplatePort
	commandTemplatePort       port.ICommandTemplatePort
	variableTemplatePort      port.IVariableTemplatePort
}

func (g GetPipelineTemplateUsecase) GetPipelineTemplate(ctx context.Context, buildTool string) (*response.PipelineTemplateResponse, error) {
	pipelineTemplates, err := g.pipelineTemplatePort.GetPipelineTemplateByBuildTools(ctx, buildTool)
	if err != nil {
		return nil, err
	}

	pipelineStageTemplates, err := g.pipelineStageTemplatePort.GetPipelineStageTemplateByTemplateID(ctx, pipelineTemplates.ID)
	if err != nil {
		return nil, err
	}
	var stageTemplateIDs []int64
	for _, pipelineStageTemplate := range pipelineStageTemplates {
		stageTemplateIDs = append(stageTemplateIDs, pipelineStageTemplate.StageID)
	}

	stageTemplates, err := g.stageTemplatePort.GetStageTemplateByIds(ctx, stageTemplateIDs)
	if err != nil {
		return nil, err
	}
	var stageIDs []int64
	for _, stageTemplate := range stageTemplates {
		stageIDs = append(stageIDs, stageTemplate.ID)
	}

	commandTemplates, err := g.commandTemplatePort.GetCommandTemplateByStageIDs(ctx, stageIDs)
	if err != nil {
		return nil, err
	}
	mapCommandTemplates := make(map[int64][]*entity.CommandTemplateEntity)
	for _, commandTemplate := range commandTemplates {
		mapCommandTemplates[commandTemplate.StageID] = append(mapCommandTemplates[commandTemplate.StageID], commandTemplate)
	}

	variableTemplates, err := g.variableTemplatePort.GetVariableTemplateByStageIDs(ctx, stageIDs)
	if err != nil {
		return nil, err
	}
	mapVariableTemplates := make(map[int64][]*entity.VariableTemplateEntity)
	for _, variableTemplate := range variableTemplates {
		mapVariableTemplates[variableTemplate.StageID] = append(mapVariableTemplates[variableTemplate.StageID], variableTemplate)
	}
	for _, stageTemplate := range stageTemplates {
		stageTemplate.Commands = mapCommandTemplates[stageTemplate.ID]
		stageTemplate.Variables = mapVariableTemplates[stageTemplate.ID]
	}
	pipelineTemplates.Stages = stageTemplates
	return response.ToPipelineTemplateResponse(pipelineTemplates), nil
}

func NewGetPipelineTemplateUsecase(pipelineTemplatePort port.IPipelineTemplatePort,
	pipelineStageTemplatePort port.IPipelineStageTemplatePort,
	stageTemplatePort port.IStageTemplatePort,
	commandTemplatePort port.ICommandTemplatePort,
	variableTemplatePort port.IVariableTemplatePort) IGetPipelineTemplateUsecase {
	return &GetPipelineTemplateUsecase{
		pipelineTemplatePort:      pipelineTemplatePort,
		pipelineStageTemplatePort: pipelineStageTemplatePort,
		stageTemplatePort:         stageTemplatePort,
		commandTemplatePort:       commandTemplatePort,
		variableTemplatePort:      variableTemplatePort,
	}
}
