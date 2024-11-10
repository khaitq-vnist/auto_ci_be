package client

import (
	"context"
	"errors"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/client"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/client/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/strategy"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
)

const (
	PathToGetUser = "/user"
)

type GithubProviderClient struct {
	httpClient client.ContextualHttpClient
	props      *properties.GithubProperties
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
