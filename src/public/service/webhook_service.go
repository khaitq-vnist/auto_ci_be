package service

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type IWebhookService interface {
	HandleUploadFileLogs(ctx context.Context, projectID, pipelineID, executionID int64) error
}
type WebhookService struct {
	uploadLogWebhookUseCase usecase.IUploadLogWebhookUseCase
}

func (w WebhookService) HandleUploadFileLogs(ctx context.Context, projectID, pipelineID, executionID int64) error {
	return w.uploadLogWebhookUseCase.FireEventUploadFileLogs(ctx, projectID, pipelineID, executionID)
}

func NewWebhookService(uploadLogWebhookUseCase usecase.IUploadLogWebhookUseCase) IWebhookService {
	return &WebhookService{
		uploadLogWebhookUseCase: uploadLogWebhookUseCase,
	}
}
