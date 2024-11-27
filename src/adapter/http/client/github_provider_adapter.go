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
	response2 "github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
)

const (
	PathToGetUser     = "/user"
	PathToGetRepos    = "/users/%s/repos"
	PathGetDetailRepo = "/repositories/%d"
)

type GithubProviderClient struct {
	httpClient client.ContextualHttpClient
	props      *properties.GithubProperties
}

func (g *GithubProviderClient) GetContentFromRepository(ctx *context.Context, token string, repoId int64, path string) (*response2.ThirdPartyContentResponse, error) {
	var content response.GithubContentResponse
	rsp, err := g.httpClient.Get(*ctx, g.props.BaseUrl+fmt.Sprintf("/repositories/%d/contents/%s", repoId, path), &content,
		client.WithHeader("Authorization", "Bearer "+token))
	if err != nil {
		log.Error(ctx, "Error when get content from github", err)
		return nil, err
	}
	if rsp.StatusCode != 200 || &content == nil {
		log.Error(ctx, "Error when get content from github", err)
		return nil, errors.New("error when get content from github")
	}
	return response.ToThirdPartyContentResponse(&content), nil
}

func (g *GithubProviderClient) GetRepositoryInfo(ctx *context.Context, token string, repoId int64) (*response2.ThirdPartyProviderReposResponse, error) {
	var repo response.GithubRepoInfo
	rsp, err := g.httpClient.Get(*ctx, g.props.BaseUrl+fmt.Sprintf(PathGetDetailRepo, repoId), &repo, client.WithHeader("Authorization", "Bearer "+token),
		client.WithHeader("Accept", "application/vnd.github+json"))
	if err != nil {
		log.Error(ctx, "Error when get repo info from github", err)
		return nil, err
	}
	if rsp.StatusCode != 200 || &repo == nil {
		log.Error(ctx, "Error when get repo info from github", err)
		return nil, errors.New("error when get repo info from github")
	}
	return response.ToThirdPartyProviderRepoResponse(&repo), nil
}

func (g *GithubProviderClient) GetListRepositoriesByUser(ctx *context.Context, token, username string) ([]*response2.ThirdPartyProviderReposResponse, error) {
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
	return response.ToListThirdPartyProviderRepoResponse(repos), nil
}

func (g *GithubProviderClient) GetUserInfo(ctx *context.Context, token string) (*response2.ThirdPartyProviderUserInfoResponse, error) {
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
