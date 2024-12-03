package request

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type BuddyCreateActionRequest struct {
	Name            string                        `json:"name"`
	Type            string                        `json:"type"`
	TriggerTime     string                        `json:"trigger_time" default:"ON_EVERY_EXECUTION"`
	DockerImageName string                        `json:"docker_image_name"`
	DockerImageTag  string                        `json:"docker_image_tag"`
	ExecuteCommands []string                      `json:"execute_commands"`
	SetupCommands   []string                      `json:"setup_commands"`
	CachedDirs      []string                      `json:"cached_dirs"`
	Variables       []*BuddyActionVariableRequest `json:"variables"`
	Shell           string                        `json:"shell" default:"BASH"`
	CacheBaseImage  bool                          `json:"cacheBaseImage"`
	Services        []*BuddyServiceRequest        `json:"services"`
}
type BuddyActionVariableRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type BuddyServiceRequest struct {
	Type       string                        `json:"type"`
	Version    string                        `json:"version"`
	Connection BuddyServiceConnectionRequest `json:"connection"`
}
type BuddyServiceConnectionRequest struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func ToBuddyActionRequest(action *entity.ActionEntity) *BuddyCreateActionRequest {
	variables := make([]*BuddyActionVariableRequest, 0)
	for _, variable := range action.Variables {
		variables = append(variables, &BuddyActionVariableRequest{
			Key:   variable.Key,
			Value: variable.Value,
		})
	}
	services := make([]*BuddyServiceRequest, 0)
	for _, service := range action.Services {
		services = append(services, &BuddyServiceRequest{
			Type:    service.Type,
			Version: service.Version,
			Connection: BuddyServiceConnectionRequest{
				Host:     service.Connection.Host,
				DB:       service.Connection.DB,
				Port:     service.Connection.Port,
				User:     service.Connection.User,
				Password: service.Connection.Password,
			},
		})
	}
	return &BuddyCreateActionRequest{
		Name:            action.Name,
		Type:            action.Type,
		TriggerTime:     action.TriggerTime,
		DockerImageName: action.DockerImageName,
		DockerImageTag:  action.DockerImageTag,
		ExecuteCommands: action.ExecuteCommands,
		SetupCommands:   action.SetupCommands,
		CachedDirs:      action.CachedDirs,
		Variables:       variables,
		Shell:           action.Shell,
		CacheBaseImage:  action.CacheBaseImage,
		Services:        services,
	}
}
