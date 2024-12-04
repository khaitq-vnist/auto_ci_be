package client

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/client"
	request2 "github.com/khaitq-vnist/auto_ci_be/adapter/http/client/dto/request"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/client/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

const (
	DefaultWorkspace = "testworkspace-10"
)

var (
	CreateNewPipelinePath = "/workspaces/%s/projects/%s/pipelines"
	CreateNewActionPath   = "/workspaces/%s/projects/%s/pipelines/%d/actions"
	GetListPipelinePath   = "/workspaces/%s/projects/%s/pipelines"
)

type ThirdPartyToolAdapter struct {
	httpClient client.ContextualHttpClient
	props      *properties.BuddyProperties
}

func (t ThirdPartyToolAdapter) GetListPipeline(ctx context.Context, project string) ([]*entity.PipelineEntity, error) {
	client := resty.New()

	var resp response.BuddyPipelineListResponse
	rsp, err := client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+t.props.AccessToken).
		SetResult(&resp).
		Get(t.props.BaseUrl + fmt.Sprintf(GetListPipelinePath, DefaultWorkspace, project))

	if err != nil {
		log.Error(ctx, "Error when getting list of pipelines:", err)
		return nil, err
	}

	if rsp.StatusCode() != 200 {
		log.Error(ctx, "Get list pipeline failed with status:", rsp.StatusCode())
		return nil, fmt.Errorf("failed to get pipelines, status code: %d", rsp.StatusCode())
	}

	return response.ToListPipelineEntities(&resp), nil
}

func (t ThirdPartyToolAdapter) CreateNewAction(ctx context.Context, project string, pipelineID int64, action *entity.ActionEntity) (*entity.ActionEntity, error) {
	client := resty.New()

	request := request2.ToBuddyActionRequest(action)
	var resp response.BuddyCreateActionResponse

	rsp, err := client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+t.props.AccessToken).
		SetBody(request).
		SetResult(&resp).
		Post(t.props.BaseUrl + fmt.Sprintf(CreateNewActionPath, DefaultWorkspace, project, pipelineID))

	if err != nil {
		log.Error(ctx, "Error when creating a new action:", err)
		return nil, err
	}

	if rsp.StatusCode() != 201 {
		log.Error(ctx, "Create new action failed with status:", rsp.StatusCode())
		return nil, fmt.Errorf("failed to create action, status code: %d", rsp.StatusCode())
	}

	return response.ToActionEntity(&resp), nil
}

func (t ThirdPartyToolAdapter) CreateNewPipeline(ctx context.Context, project string, pipeline *entity.PipelineEntity) (*entity.PipelineEntity, error) {
	request := request2.ToBuddyPipelineRequest(pipeline)
	requestBody, _ := request.ToJson()

	// Create a new Resty client
	client := resty.New()

	// Define the API URL
	url := t.props.BaseUrl + fmt.Sprintf(CreateNewPipelinePath, DefaultWorkspace, project)

	// Perform the POST request
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+t.props.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		SetResult(&response.BuddyPipelineResponse{}).
		Post(url)

	if err != nil {
		log.Error(ctx, "Error when creating new pipeline:", err)
		return nil, err
	}

	// Check the response status code
	if resp.StatusCode() != 201 {
		log.Error(ctx, "Create new pipeline failed:", resp.Status(), resp.String())
		return nil, fmt.Errorf("failed to create new pipeline: %s", resp.Status())
	}

	// Parse the response
	respBody := resp.Result().(*response.BuddyPipelineResponse)

	// Convert to the expected entity
	return response.ToPipelineEntity(respBody), nil
}

func NewThirdPartyToolAdapter(httpClient client.ContextualHttpClient, props *properties.BuddyProperties) port.IThirdPartyToolPort {
	return &ThirdPartyToolAdapter{
		httpClient: httpClient,
		props:      props,
	}
}
