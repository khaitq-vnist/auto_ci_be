package client

import (
	"context"
	"errors"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/strategy"
	"github.com/khaitq-vnist/auto_ci_be/core/constant"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ThirdPartyProviderAdapter struct {
	strategies map[string]strategy.IThirdPartyStrategy
}

func (t *ThirdPartyProviderAdapter) GetListBranches(ctx *context.Context, provider string, username, token, repo string) ([]*response.ThirdPartyBranchResponse, error) {
	partyStrategy := t.getStrategy(provider)
	if partyStrategy == nil {
		log.Error(ctx, "Provider not found", nil)
		return nil, errors.New("provider not found")
	}
	branches, err := partyStrategy.GetListBranches(ctx, username, token, repo)
	if err != nil {
		log.Error(ctx, "Error when get branches from third party provider", err)
		return nil, err
	}
	return branches, nil
}

func (t *ThirdPartyProviderAdapter) GetContentFromRepository(ctx *context.Context, provider, username, token, repo, path string) ([]*response.ThirdPartyContentResponse, error) {
	partyStrategy := t.getStrategy(provider)
	if partyStrategy == nil {
		log.Error(ctx, "Provider not found", nil)
		return nil, errors.New("provider not found")
	}
	contents, err := partyStrategy.GetContentFromRepository(ctx, username, token, repo, path)
	if err != nil {
		log.Error(ctx, "Error when get content from third party provider", err)
		return nil, err
	}
	return contents, nil
}

func (t *ThirdPartyProviderAdapter) GetRepositoryInfo(ctx *context.Context, provider string, token string, repoId int64) (*response.ThirdPartyProviderReposResponse, error) {
	partyStrategy := t.getStrategy(provider)
	if partyStrategy == nil {
		log.Error(ctx, "Provider not found", nil)
		return nil, errors.New("provider not found")
	}
	repo, err := partyStrategy.GetRepositoryInfo(ctx, token, repoId)
	if err != nil {
		log.Error(ctx, "Error when get repository info from third party provider", err)
		return nil, err
	}
	return repo, nil
}

func (t *ThirdPartyProviderAdapter) GetListRepositoriesByUser(ctx *context.Context, provider string, token, username string) ([]*response.ThirdPartyProviderReposResponse, error) {
	partyStrategy := t.getStrategy(provider)
	if partyStrategy == nil {
		log.Error(ctx, "Provider not found", nil)
		return nil, errors.New("provider not found")
	}
	repos, err := partyStrategy.GetListRepositoriesByUser(ctx, token, username)
	if err != nil {
		log.Error(ctx, "Error when get list repositories from third party provider", err)
		return nil, err
	}
	return repos, nil
}

func (t *ThirdPartyProviderAdapter) GetUserInfo(ctx context.Context, provider string, token string) (*response.ThirdPartyProviderUserInfoResponse, error) {
	partyStrategy := t.getStrategy(provider)
	if partyStrategy == nil {
		log.Error(ctx, "Provider not found", nil)
		return nil, errors.New("provider not found")
	}
	userInfo, err := partyStrategy.GetUserInfo(ctx, token)
	if err != nil {
		log.Error(ctx, "Error when get user info from third party provider", err)
		return nil, err
	}
	return userInfo, nil
}

func NewThirdPartyProviderAdapter(gitHubStrategy strategy.IThirdPartyStrategy,
	gitLabStrategy strategy.IThirdPartyStrategy) port.IThirdPartyProviderPort {
	strategies := map[string]strategy.IThirdPartyStrategy{
		constant.GITHUB_PROVIDER: gitHubStrategy,
		constant.GITLAB_PROVIDER: gitLabStrategy,
	}
	return &ThirdPartyProviderAdapter{
		strategies: strategies,
	}
}

func (t *ThirdPartyProviderAdapter) getStrategy(provider string) strategy.IThirdPartyStrategy {
	partyStrategy, ok := t.strategies[provider]
	if !ok {
		return nil
	}
	return partyStrategy
}
