package client

import (
	"context"
	"errors"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/strategy"
	"github.com/khaitq-vnist/auto_ci_be/core/constanst"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type ThirdPartyProviderAdapter struct {
	strategies map[string]strategy.IThirdPartyStrategy
}

func (t *ThirdPartyProviderAdapter) GetListRepositoriesByUser(ctx *context.Context, provider string, token, username string) ([]*dto.ThirdPartyProviderReposResponse, error) {
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

func (t *ThirdPartyProviderAdapter) GetUserInfo(ctx *context.Context, provider string, token string) (*dto.ThirdPartyProviderUserInfoResponse, error) {
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
		constanst.GITHUB_PROVIDER: gitHubStrategy,
		constanst.GITLAB_PROVIDER: gitLabStrategy,
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
