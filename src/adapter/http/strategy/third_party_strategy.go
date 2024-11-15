package strategy

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
)

type IThirdPartyStrategy interface {
	GetUserInfo(ctx *context.Context, token string) (*response.ThirdPartyProviderUserInfoResponse, error)
	GetListRepositoriesByUser(ctx *context.Context, token string, username string) ([]*response.ThirdPartyProviderReposResponse, error)
	GetRepositoryInfo(ctx *context.Context, token string, repoId int64) (*response.ThirdPartyProviderReposResponse, error)
}
