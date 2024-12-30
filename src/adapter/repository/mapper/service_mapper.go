package mapper

import (
	"encoding/json"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type ServiceConnectionJson struct {
	Host     string `json:"host,omitempty"`
	DB       string `json:"db,omitempty"`
	Port     string `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

func ToListServiceEntityMapper(models []*model.ServiceModel) []*entity.ServiceEntity {
	entities := make([]*entity.ServiceEntity, 0)
	if models == nil {
		return entities
	}
	for _, serviceModel := range models {
		entities = append(entities, ToServiceEntityMapper(serviceModel))
	}
	return entities
}
func ToServiceEntityMapper(model *model.ServiceModel) *entity.ServiceEntity {
	if model == nil {
		return nil
	}
	var connection ServiceConnectionJson

	if model.Connection != nil {
		err := json.Unmarshal(model.Connection, &connection)
		if err != nil {
			connection = ServiceConnectionJson{}
		}
	}
	return &entity.ServiceEntity{
		BaseEntity: entity.BaseEntity{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix(),
		},
		Type:    model.Type,
		Version: model.Version,
		Connection: entity.ServiceConnectionEntity{
			Host:     connection.Host,
			DB:       connection.DB,
			Port:     connection.Port,
			User:     connection.User,
			Password: connection.Password,
		},
	}
}
