package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"sort"
)

type IAnalyzeSourceCodeUsecase interface {
	Analyze(ctx context.Context, userID, projectId int64) (*dto.AnalyzeSourceCodeDTO, error)
}
type AnalyzeSourceCodeUsecase struct {
	thirdPartyProviderPort port.IThirdPartyProviderPort
	getProjectUseCase      IGetProjectUseCase
	getIntegrationUseCase  IGetIntegrationUseCase
	encryptUseCase         IEncryptUseCase
}

func (a AnalyzeSourceCodeUsecase) Analyze(ctx context.Context, userID, projectId int64) (*dto.AnalyzeSourceCodeDTO, error) {
	project, err := a.getProjectUseCase.GetProjectById(ctx, projectId)
	if err != nil {
		return nil, err
	}
	integration, err := a.getIntegrationUseCase.GetIntegrationByIdAndUserId(ctx, project.IntegrationId, 1)
	if err != nil {
		return nil, err
	}
	if integration.ProviderName == "GitHub" {
		integration.ProviderName = "GITHUB"
	}
	project.ProviderUsername = integration.ProviderUsername
	decryptToken, err := a.encryptUseCase.DecryptToken(&ctx, integration.AccessToken)
	tools, err := a.detectTools(ctx, project, integration.ProviderName, decryptToken, "")
	if err != nil {
		log.Error(ctx, "AnalyzeSourceCodeUsecase", "Analyze", err)
		return nil, err
	}
	return &dto.AnalyzeSourceCodeDTO{
		Tools: tools,
	}, nil
}

func NewAnalyzeSourceCodeUsecase(
	thirdPartyProviderPort port.IThirdPartyProviderPort,
	getProjectUseCase IGetProjectUseCase,
	getIntegrationUseCase IGetIntegrationUseCase,
	encryptUseCase IEncryptUseCase) IAnalyzeSourceCodeUsecase {
	return &AnalyzeSourceCodeUsecase{
		thirdPartyProviderPort: thirdPartyProviderPort,
		getProjectUseCase:      getProjectUseCase,
		getIntegrationUseCase:  getIntegrationUseCase,
		encryptUseCase:         encryptUseCase,
	}
}
func (a AnalyzeSourceCodeUsecase) detectTools(ctx context.Context, project *entity.ProjectEntity, provider, token, path string) (map[string]bool, error) {
	tools := map[string]bool{
		"Maven":   false,
		"Gradle":  false,
		"Node.js": false,
		"Go":      false,
		"Python":  false,
		"Ruby":    false,
	}
	contents, err := a.thirdPartyProviderPort.GetContentFromRepository(&ctx, provider, project.ProviderUsername, token, project.Name, path)
	sort.Slice(contents, func(i, j int) bool {
		return contents[i].Type > contents[j].Type
	})
	if err != nil {
		return nil, err
	}
	identify := false
	for _, item := range contents {
		if item.Type == "file" {
			if item.Name == "pom.xml" {
				tools["Maven"] = true
				identify = true
			}
			if item.Name == "build.gradle" {
				tools["Gradle"] = true
				identify = true
			}
			if item.Name == "package.json" {
				tools["Node.js"] = true
				identify = true
			}
			if item.Name == "go.mod" {
				tools["Go"] = true
				identify = true
			}
			if item.Name == "requirements.txt" {
				tools["Python"] = true
				identify = true
			}
			if item.Name == "Gemfile" {
				tools["Ruby"] = true
				identify = true
			}
			if identify {
				return tools, nil
			}
		} else if item.Type == "dir" {
			subTools, err := a.detectTools(ctx, project, provider, token, item.Path)
			if err != nil {
				return nil, err
			}
			for key, value := range subTools {
				tools[key] = tools[key] || value
			}
		}
	}
	return tools, nil
}
