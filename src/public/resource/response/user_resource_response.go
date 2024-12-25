package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type UserResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func FromEntityToUserResponse(entity *entity.UserEntity) *UserResponse {
	if entity == nil {
		return nil
	}
	return &UserResponse{
		ID:        entity.ID,
		Email:     entity.Email,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
