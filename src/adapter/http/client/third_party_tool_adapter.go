package client

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/client"
	request2 "github.com/khaitq-vnist/auto_ci_be/adapter/http/client/dto/request"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/client/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

const (
	DefaultWorkspace = "test-1"
)

var (
	CreateNewPipelinePath = "/workspaces/%s/projects/%s/pipelines"
	CreateNewActionPath   = "/workspaces/%s/projects/%s/pipelines/%d/actions"
)

type ThirdPartyToolAdapter struct {
	httpClient client.ContextualHttpClient
	props      *properties.BuddyProperties
}

func (t ThirdPartyToolAdapter) CreateNewAction(ctx context.Context, project string, pipelineID int64, action *entity.ActionEntity) (*entity.ActionEntity, error) {
	request := request2.ToBuddyActionRequest(action)
	var resp response.BuddyCreateActionResponse
	rsp, err := t.httpClient.Post(ctx, t.props.BaseUrl+fmt.Sprint(CreateNewActionPath, DefaultWorkspace, project, pipelineID), request, resp,
		client.WithHeader("Authorization", "Bearer "+t.props.AccessToken))
	if err != nil {
		log.Error(ctx, "Error when create new action", err)
		return nil, err
	}
	if rsp.StatusCode != 201 {
		log.Error(ctx, "Create new action failed")
		return nil, err
	}
	return response.ToActionEntity(&resp), nil
}

func (t ThirdPartyToolAdapter) CreateNewPipeline(ctx context.Context, project string, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error) {
	request := request2.ToBuddyPipelineRequest(pipeline)
	var resp response.BuddyPipelineResponse
	rsp, err := t.httpClient.Post(ctx, t.props.BaseUrl+fmt.Sprintf(CreateNewPipelinePath, DefaultWorkspace, project), request, resp,
		client.WithHeader("Authorization", "Bearer "+t.props.AccessToken))
	if err != nil {
		log.Error(ctx, "Error when create new pipeline", err)
		return nil, err
	}
	if rsp.StatusCode != 201 {
		log.Error(ctx, "Create new pipeline failed")
		return nil, err
	}
	return response.ToPipelineEntity(&resp), nil
}

func NewThirdPartyToolAdapter(httpClient client.ContextualHttpClient, props *properties.BuddyProperties) port.IThirdPartyToolPort {
	return &ThirdPartyToolAdapter{
		httpClient: httpClient,
		props:      props,
	}
}
