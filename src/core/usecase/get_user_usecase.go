package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IGetUserUseCase interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}
type GetUserUseCase struct {
	userPort port.IUserPort
}

func (g GetUserUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	user, err := g.userPort.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(ctx, "Get user failed: ", err)
		return nil, err
	}
	return user, nil
}

func NewGetUserUseCase(userPort port.IUserPort) IGetUserUseCase {
	return &GetUserUseCase{userPort}
}
