package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

// PipelineTemplateResponse is a response struct for pipeline template
type PipelineTemplateResponse struct {
	ID          int64                    `json:"id"`
	Name        string                   `json:"name"`
	BuildTool   string                   `json:"build_tool"`
	Description string                   `json:"description"`
	Stages      []*StageTemplateResponse `json:"stages"`
}

func ToPipelineTemplateResponse(pipeline *entity.PipelineTemplateEntity) *PipelineTemplateResponse {
	return &PipelineTemplateResponse{
		ID:          pipeline.ID,
		Name:        pipeline.Name,
		BuildTool:   pipeline.BuildTool,
		Description: pipeline.Description,
		Stages:      ToListStageTemplateResponse(pipeline.Stages),
	}
}

type StageTemplateResponse struct {
	ID             int64                      `json:"id"`
	Name           string                     `json:"name"`
	Type           string                     `json:"type"`
	DockerImage    string                     `json:"docker_image"`
	DockerImageTag string                     `json:"docker_image_tag"`
	Commands       []*CommandTemplateResponse `json:"commands"`
	Variables      []*VariableResponse        `json:"variables"`
	Services       []*ServiceTemplateResponse `json:"services"`
}

func ToListStageTemplateResponse(stages []*entity.StageTemplateEntity) []*StageTemplateResponse {
	var stageResponses []*StageTemplateResponse
	for _, stage := range stages {
		stageResponses = append(stageResponses, ToStageTemplateResponse(stage))
	}
	return stageResponses
}
func ToStageTemplateResponse(stage *entity.StageTemplateEntity) *StageTemplateResponse {
	return &StageTemplateResponse{
		ID:             stage.ID,
		Name:           stage.Name,
		Type:           stage.Type,
		DockerImage:    stage.DockerImage,
		DockerImageTag: stage.DockerImageTag,
		Commands:       ToCommandTemplateResponses(stage.Commands),
		Variables:      ToVariableResponses(stage.Variables),
	}
}

type CommandTemplateResponse struct {
	ID      int64  `json:"id"`
	Command string `json:"command"`
}

func ToCommandTemplateResponses(commands []*entity.CommandTemplateEntity) []*CommandTemplateResponse {
	var commandResponses []*CommandTemplateResponse
	for _, command := range commands {
		commandResponses = append(commandResponses, &CommandTemplateResponse{
			ID:      command.ID,
			Command: command.Command,
		})
	}
	return commandResponses
}

type VariableResponse struct {
	ID    int64  `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ToVariableResponses(variables []*entity.VariableTemplateEntity) []*VariableResponse {
	var variableResponses []*VariableResponse
	for _, variable := range variables {
		variableResponses = append(variableResponses, &VariableResponse{
			ID:    variable.ID,
			Key:   variable.Key,
			Value: variable.Value,
		})
	}
	return variableResponses
}

type PipelineSettingsResponse struct {
	ID        int64                   `json:"id"`
	Status    string                  `json:"status"`
	Trigger   PipelineTriggerResponse `json:"trigger"`
	CodeScope CodeScopeResponse       `json:"code_scope"`
}
type PipelineTriggerResponse struct {
	ID          int64  `json:"id"`
	TriggerType string `json:"trigger_type"`
	TriggerName string `json:"trigger_name"`
}

type CodeScopeResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
type ServiceTemplateResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Version  string `json:"version"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
