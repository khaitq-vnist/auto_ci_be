package request

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=64"`
	Name     string `json:"name" validate:"required,gte=6,lte=64"`
}

func FromCreateUserReqToUserEntity(req *CreateUserRequest) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
}
