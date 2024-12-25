package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/exception"
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
	if err != nil && err.Error() != common.ErrRecordNotFound {
		log.Error(ctx, "Get user failed: ", err)
		return nil, err
	}
	if existedUser != nil {
		log.Error(ctx, fmt.Sprintf("Email %s existed", user.Email))
		return nil, errors.New(common.ExistedEmailMessage)
	}
	tx := c.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if err = c.databaseTransactionUsecase.Rollback(tx); err != nil {
			log.Error(ctx, "Commit transaction failed: ", err)
		} else {
			log.Info(ctx, "Commit transaction success")
		}
	}()
	hashPassword, err := common.HashPassword(user.Password)
	if err != nil {
		log.Error(ctx, "Hash password failed: ", err)
		return nil, err
	}
	user.Password = hashPassword
	user, err = c.userPort.SaveUser(ctx, tx, user)
	if err != nil {
		log.Error(ctx, "Save user failed: ", err)
		return nil, err
	}
	if err = c.databaseTransactionUsecase.Commit(tx); err != nil {
		log.Error(ctx, "Commit transaction failed: ", err)
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func NewCreateUserUseCase(userPort port.IUserPort, databaseTransactionUsecase IDatabaseTransactionUsecase, getUserUseCase IGetUserUseCase) ICreateUserUseCase {
	return &CreateUserUseCase{userPort: userPort, databaseTransactionUsecase: databaseTransactionUsecase, getUserUseCase: getUserUseCase}
}
