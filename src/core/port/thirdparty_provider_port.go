package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
)

type IThirdPartyProviderPort interface {
	GetUserInfo(ctx context.Context, provider string, token string) (*response.ThirdPartyProviderUserInfoResponse, error)
	GetListRepositoriesByUser(ctx *context.Context, provider string, token, username string) ([]*response.ThirdPartyProviderReposResponse, error)
	GetRepositoryInfo(ctx *context.Context, provider string, token string, repoId int64) (*response.ThirdPartyProviderReposResponse, error)

	GetContentFromRepository(ctx *context.Context, provider string, username, token, repo string, path string) ([]*response.ThirdPartyContentResponse, error)

	GetListBranches(ctx *context.Context, provider string, username, token, repo string) ([]*response.ThirdPartyBranchResponse, error)
}
