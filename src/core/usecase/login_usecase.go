package usecase

import (
	"context"
	"errors"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/properties"
)

type ILoginUseCase interface {
	Login(ctx context.Context, email string, password string) (*response.LoginResponseDto, error)
}
type LoginUseCase struct {
	jwtProps       *properties.TokenProperties
	getUserUseCase IGetUserUseCase
}

func (l LoginUseCase) Login(ctx context.Context, email string, password string) (*response.LoginResponseDto, error) {
	exitedEmail, err := l.getUserUseCase.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(ctx, "GetUserByEmail error: ", err)
		return nil, err
	}
	ok := common.ComparePassword(exitedEmail.Password, password)
	if !ok {
		log.Error(ctx, "Password is incorrect")
		return nil, errors.New(common.ErrInvalidPassword)
	}
	token, err := common.GenerateToken(exitedEmail, l.jwtProps)
	if err != nil {
		log.Error(ctx, "Generate token error: ", err)
		return nil, err
	}
	var logRsp response.LoginResponseDto
	logRsp.AccessToken = token
	return &logRsp, nil
}

func NewLoginUseCase(jwtProps *properties.TokenProperties, getUserUseCase IGetUserUseCase) ILoginUseCase {
	return &LoginUseCase{
		jwtProps:       jwtProps,
		getUserUseCase: getUserUseCase,
	}
}
