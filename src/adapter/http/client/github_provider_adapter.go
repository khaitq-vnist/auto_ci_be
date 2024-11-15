package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/client"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/client/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/strategy"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
)

const (
	PathToGetUser  = "/user"
	PathToGetRepos = "/users/%s/repos"
)

type GithubProviderClient struct {
	httpClient client.ContextualHttpClient
	props      *properties.GithubProperties
}

func (g *GithubProviderClient) GetListRepositoriesByUser(ctx *context.Context, token, username string) ([]*dto.ThirdPartyProviderReposResponse, error) {
	var repos []*response.GithubRepoInfo
	rsp, err := g.httpClient.Get(*ctx, g.props.BaseUrl+fmt.Sprintf(PathToGetRepos, username), &repos,
		client.WithHeader("Authorization", "Bearer "+token),
		client.WithHeader("Accept", "application/vnd.github+json"),
		client.WithHeader("X-GitHub-Api-Version", "2022-11-28"))
	if err != nil {
		log.Error(ctx, "Error when get list repos from github", err)
		return nil, err
	}
	if rsp.StatusCode != 200 || &repos == nil {
		log.Error(ctx, "Error when get list repos from github", err)
		return nil, errors.New("error when get list repos from github")
	}
	return response.ToThirdPartyProviderProjectResponse(repos), nil
}

func (g *GithubProviderClient) GetUserInfo(ctx *context.Context, token string) (*dto.ThirdPartyProviderUserInfoResponse, error) {
	var userInfoResp response.GithubUserInfoResponse
	rsp, err := g.httpClient.Get(*ctx, g.props.BaseUrl+PathToGetUser, &userInfoResp,
		client.WithHeader("Authorization", "Bearer "+token))
	if err != nil {
		log.Error(ctx, "Error when get user info from github", err)
		return nil, err
	}
	if rsp.StatusCode != 200 || &userInfoResp == nil {
		log.Error(ctx, "Error when get user info from github", err)
		return nil, errors.New("error when get user info from github")
	}
	return userInfoResp.ToUserInfo(), nil
}

func NewGithubProviderClient(httpClient client.ContextualHttpClient, props *properties.GithubProperties) strategy.IThirdPartyStrategy {
	return &GithubProviderClient{
		httpClient: httpClient,
		props:      props,
	}
}
