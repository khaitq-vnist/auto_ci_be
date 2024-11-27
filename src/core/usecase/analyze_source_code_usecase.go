package usecase

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IAnalyzeSourceCodeUsecase interface {
	Analyze(ctx context.Context, userID, projectId int64) (*dto.AnalyzeSourceCodeDTO, error)
}
type AnalyzeSourceCodeUsecase struct {
	thirdPartyProviderPort port.IThirdPartyProviderPort
	getProjectUseCase      IGetProjectUseCase
	getIntegrationUseCase  IGetIntegrationUseCase
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
	decryptToken, err := a.encryptUseCase.DecryptToken(&ctx, integration.AccessToken)
	if err != nil {
		return nil, err
	}
}

func NewAnalyzeSourceCodeUsecase(thirdPartyProviderPort port.IThirdPartyProviderPort, getProjectUseCase IGetProjectUseCase) IAnalyzeSourceCodeUsecase {
	return &AnalyzeSourceCodeUsecase{
		thirdPartyProviderPort: thirdPartyProviderPort,
		getProjectUseCase:      getProjectUseCase,
	}
}
func (a AnalyzeSourceCodeUsecase) detectTools(ctx context.Context, project *entity.ProjectEntity, provider, token string) ([]string, error) {
	tools := map[string]bool{
		"Maven":   false,
		"Gradle":  false,
		"Node.js": false,
		"Go":      false,
		"Python":  false,
		"Ruby":    false,
		"Docker":  false,
	}
	content, err := a.thirdPartyProviderPort.GetContentFromRepository(ctx, provider, token, project.R, "pom.xml")
	return tools, nil
}
