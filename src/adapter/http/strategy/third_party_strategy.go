package strategy

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
)

type IThirdPartyStrategy interface {
	GetUserInfo(ctx *context.Context, token string) (*dto.ThirdPartyProviderUserInfoResponse, error)
	GetListRepositoriesByUser(ctx *context.Context, token string, username string) ([]*dto.ThirdPartyProviderReposResponse, error)
}
