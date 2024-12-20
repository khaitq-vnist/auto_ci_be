package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ICreateUserUseCase interface {
	CreateUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
}
type CreateUserUseCase struct {
	userPort                   port.IUserPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
	getUserUseCase             IGetUserUseCase
}

func (c CreateUserUseCase) CreateUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {
	existedUser, err := c.getUserUseCase.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Error(ctx, "Get user failed: ", err)
		return nil, err
	}
	if existedUser != nil {
		log.Error(ctx, fmt.Sprintf("Email %s existed", user.Email))
		return nil, errors.New(common.ExistedEmailMessage)
	}
	tx := c.databaseTransactionUsecase.StartTransaction()
}

func NewCreateUserUseCase(userPort port.IUserPort, databaseTransactionUsecase IDatabaseTransactionUsecase) ICreateUserUseCase {
	return &CreateUserUseCase{userPort, databaseTransactionUsecase}
}
