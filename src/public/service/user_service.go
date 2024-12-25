package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/request"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IUserService interface {
	CreateUser(ctx context.Context, req *request.CreateUserRequest) (*entity.UserEntity, error)
	LoginUser(ctx context.Context, email, password string) (*response.LoginResponseDto, error)
}
type UserService struct {
	createUserUseCase usecase.ICreateUserUseCase
	loginUseCase      usecase.ILoginUseCase
}

func (u UserService) LoginUser(ctx context.Context, email, password string) (*response.LoginResponseDto, error) {
	return u.loginUseCase.Login(ctx, email, password)
}

func (u UserService) CreateUser(ctx context.Context, req *request.CreateUserRequest) (*entity.UserEntity, error) {
	userEntity := request.FromCreateUserReqToUserEntity(req)
	return u.createUserUseCase.CreateUser(ctx, userEntity)
}
func NewUserService(createUserUseCase usecase.ICreateUserUseCase, loginUseCase usecase.ILoginUseCase) IUserService {
	return &UserService{
		createUserUseCase: createUserUseCase,
		loginUseCase:      loginUseCase,
	}
}
