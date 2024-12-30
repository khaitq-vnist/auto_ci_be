package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type BuddyCreateActionResponse struct {
	ID              int64                          `json:"id"`
	Name            string                         `json:"name"`
	Type            string                         `json:"type"`
	TriggerTime     string                         `json:"trigger_time" default:"ON_EVERY_EXECUTION"`
	DockerImageName string                         `json:"docker_image_name"`
	DockerImageTag  string                         `json:"docker_image_tag"`
	ExecuteCommands []string                       `json:"execute_commands"`
	SetupCommands   []string                       `json:"setup_commands"`
	CachedDirs      []string                       `json:"cached_dirs"`
	Variables       []*BuddyActionVariableResponse `json:"variables"`
	Shell           string                         `json:"shell" default:"BASH"`
	CacheBaseImage  bool                           `json:"cacheBaseImage"`
	Services        []*BuddyServiceResponse        `json:"services"`
}
type BuddyActionVariableResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type BuddyServiceResponse struct {
	Type       string                         `json:"type"`
	Version    string                         `json:"version"`
	Connection BuddyServiceConnectionResponse `json:"connection"`
}
type BuddyServiceConnectionResponse struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func ToActionEntity(buddyRsp *BuddyCreateActionResponse) *entity.ActionEntity {
	variables := make([]*entity.ActionVariableEntity, 0)
	for _, variable := range buddyRsp.Variables {
		variables = append(variables, &entity.ActionVariableEntity{
			Key:   variable.Key,
			Value: variable.Value,
		})
	}
	services := make([]*entity.ServiceEntity, 0)
	for _, service := range buddyRsp.Services {
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
	return &entity.ActionEntity{
		ID:              buddyRsp.ID,
		Name:            buddyRsp.Name,
		Type:            buddyRsp.Type,
		TriggerTime:     buddyRsp.TriggerTime,
		DockerImageName: buddyRsp.DockerImageName,
		DockerImageTag:  buddyRsp.DockerImageTag,
		ExecuteCommands: buddyRsp.ExecuteCommands,
		SetupCommands:   buddyRsp.SetupCommands,
		CachedDirs:      buddyRsp.CachedDirs,
		Variables:       variables,
		Shell:           buddyRsp.Shell,
		CacheBaseImage:  buddyRsp.CacheBaseImage,
		Services:        services,
	}
}
