package request

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type CreatePipelineRequest struct {
	Name   string                  `json:"name"`
	On     string                  `json:"on"`
	Refs   []string                `json:"refs"`
	Events []*EventPipelineRequest `json:"events"`
	Stages []*StagePipelineRequest `json:"stages"`
}
type EventPipelineRequest struct {
	Type string   `json:"type"`
	Refs []string `json:"refs"`
}
type StagePipelineRequest struct {
	Name            string                   `json:"name"`
	Type            string                   `json:"type"`
	TriggerTime     string                   `json:"trigger_time"`
	DockerImageName string                   `json:"docker_image_name"`
	DockerImageTag  string                   `json:"docker_image_tag"`
	ExecuteCommands []string                 `json:"execute_commands"`
	SetupCommands   []string                 `json:"setup_commands"`
	CachedDirs      []string                 `json:"cached_dirs"`
	Variables       []*ActionVariableRequest `json:"variables"`
	Shell           string                   `json:"shell"`
	CacheBaseImage  bool                     `json:"cacheBaseImage"`
	Services        []*ServiceRequest        `json:"services"`
}
type ActionVariableRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type ServiceRequest struct {
	Type       string                   `json:"type"`
	Version    string                   `json:"version"`
	Connection ServiceConnectionRequest `json:"connection"`
}
type ServiceConnectionRequest struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func ToPipelineEntity(req *CreatePipelineRequest) *entity.PipelineEntity {
	if req == nil {
		return nil
	}
	events := make([]*entity.EventEntity, 0)
	for _, event := range req.Events {
		events = append(events, &entity.EventEntity{
			Type: event.Type,
			Refs: event.Refs,
		})
	}
	stages := make([]*entity.ActionEntity, 0)
	for _, stage := range req.Stages {
		variables := make([]*entity.ActionVariableEntity, 0)
		for _, variable := range stage.Variables {
			variables = append(variables, &entity.ActionVariableEntity{
				Key:   variable.Key,
				Value: variable.Value,
			})
		}
		services := make([]*entity.ServiceEntity, 0)
		for _, service := range stage.Services {
			services = append(services, &entity.ServiceEntity{
				Type:    service.Type,
				Version: service.Version,
				Connection: entity.ServiceConnectionEntity{
					Host:     service.Connection.Host,
					DB:       service.Connection.DB,
					Port:     service.Connection.Port,
					User:     service.Connection.User,
					Password: service.Connection.Password,
				},
			})
		}
		stages = append(stages, &entity.ActionEntity{
			Name:            stage.Name,
			Type:            stage.Type,
			TriggerTime:     stage.TriggerTime,
			DockerImageName: stage.DockerImageName,
			DockerImageTag:  stage.DockerImageTag,
			ExecuteCommands: stage.ExecuteCommands,
			SetupCommands:   stage.SetupCommands,
			CachedDirs:      stage.CachedDirs,
			Variables:       variables,
			Shell:           stage.Shell,
			CacheBaseImage:  stage.CacheBaseImage,
			Services:        services,
		})
	}
	return &entity.PipelineEntity{
		Name:    req.Name,
		On:      req.On,
		Refs:    req.Refs,
		Events:  events,
		Actions: stages,
	}
}
