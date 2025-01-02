package client

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golibs-starter/golib/log"
	response2 "github.com/khaitq-vnist/auto_ci_be/adapter/http/client/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"strconv"
)

var (
	sonarCreateProjectUrl = "/api/projects/create"
	sonarGenerateTokenUrl = "/api/user_tokens/generate"
	sonarGetCoverageUrl   = "/api/measures/component"
)

type SonarClientAdapter struct {
	props *properties.SonarProperties
}

func (s SonarClientAdapter) GetCoverage(ctx context.Context, projectKey string) (float64, error) {
	client := resty.New()
	var rsp response2.CoverageResponse
	resp, err := client.R().
		SetResult(&rsp).
		SetHeader("Authorization", "Bearer "+s.props.AccessToken).
		SetQueryParams(map[string]string{
			"component":  projectKey,
			"metricKeys": "coverage",
		}).Get(s.props.BaseUrl + sonarGetCoverageUrl)
	if err != nil {
		log.Error(ctx, "Error when get coverage", err)
		return 0, err
	}
	if resp.StatusCode() != 200 {
		log.Error(ctx, "Status code is not 200")
		return 0, fmt.Errorf("status code is not 200")
	}
	var result float64
	if len(rsp.Component.Measures) > 0 {
		result, err = strconv.ParseFloat(rsp.Component.Measures[0].Value, 64)
		if err != nil {
			log.Error(ctx, "Error when parse float", err)
			return 0, err
		}
	}
	return result, nil
}

func (s SonarClientAdapter) GetAccessToken(ctx context.Context, name, projectKey string) (string, error) {
	client := resty.New()
	var rsp response2.SonarCreateTokenResponse
	resp, err := client.R().
		SetResult(&rsp).
		SetHeader("Authorization", "Bearer "+s.props.AccessToken).
		SetQueryParams(map[string]string{
			"name":       name,
			"projectKey": projectKey,
			"type":       "PROJECT_ANALYSIS_TOKEN",
		}).Post(s.props.BaseUrl + sonarGenerateTokenUrl)
	if err != nil {
		log.Error(ctx, "Error when get access token", err)
		return "", err
	}
	if resp.StatusCode() != 200 {
		log.Error(ctx, "Status code is not 200")
		return "", fmt.Errorf("status code is not 200")
	}
	return rsp.Token, nil
}

func (s SonarClientAdapter) CreateNewProject(ctx context.Context, name, project string) (*response.SonarCreateProjectResponse, error) {
	var rsp response2.SonarCreateProjectResponse
	client := resty.New()
	resp, err := client.R().
		SetResult(&rsp).
		SetHeader("Authorization", "Bearer "+s.props.AccessToken).
		SetQueryParams(map[string]string{
			"name":       name,
			"project":    project,
			"mainBranch": "master",
		}).
		Post(s.props.BaseUrl + sonarCreateProjectUrl)
	if err != nil {
		log.Error(ctx, "Error when create new project", err)
		return nil, err
	}
	if resp.StatusCode() != 200 {
		log.Error(ctx, "Status code is not 200")
		return nil, fmt.Errorf("status code is not 200")
	}
	return response2.ToSonarCreateProjectResponse(&rsp), nil
}

func NewSonarClientAdapter(props *properties.SonarProperties) port.IQualityGatePort {
	return &SonarClientAdapter{props: props}
}
