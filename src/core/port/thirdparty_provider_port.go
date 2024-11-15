package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
)

type IThirdPartyProviderPort interface {
	GetUserInfo(ctx *context.Context, provider string, token string) (*dto.ThirdPartyProviderUserInfoResponse, error)
	GetListRepositoriesByUser(ctx *context.Context, provider string, token, username string) ([]*dto.ThirdPartyProviderReposResponse, error)
}
