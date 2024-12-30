package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type ServiceResponse struct {
	ID         int64                     `json:"id"`
	Type       string                    `json:"type"`
	Version    string                    `json:"version"`
	Connection ServiceConnectionResponse `json:"connection"`
}
type ServiceConnectionResponse struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func FromListEntityToServiceResponse(entities []*entity.ServiceEntity) []*ServiceResponse {
	if entities == nil {
		return nil
	}
	responses := make([]*ServiceResponse, 0)
	for _, entity := range entities {
		responses = append(responses, FromEntityToServiceResponse(entity))
	}
	return responses
}
func FromEntityToServiceResponse(entity *entity.ServiceEntity) *ServiceResponse {
	if entity == nil {
		return nil
	}
	var connection ServiceConnectionResponse

	connection = ServiceConnectionResponse{
		Host:     entity.Connection.Host,
		DB:       entity.Connection.DB,
		Port:     entity.Connection.Port,
		User:     entity.Connection.User,
		Password: entity.Connection.Password,
	}

	return &ServiceResponse{
		ID:         entity.ID,
		Type:       entity.Type,
		Version:    entity.Version,
		Connection: connection,
	}
}
