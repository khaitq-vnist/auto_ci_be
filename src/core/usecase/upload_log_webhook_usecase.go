package usecase

import (
	"context"
	"github.com/golibs-starter/golib/log"
	message2 "github.com/khaitq-vnist/auto_ci_be/core/event/message"
)

type IUploadLogWebhookUseCase interface {
	FireEventUploadFileLogs(ctx context.Context, projectID, pipelineID, executionID int64) error
}
type UploadLogWebhookUseCase struct {
	fireEventUsecase IFireEventUsecase
}

func (u UploadLogWebhookUseCase) FireEventUploadFileLogs(ctx context.Context, projectID, pipelineID, executionID int64) error {
	message := message2.LogsEventMessage{
		UserId:      1,
		ProjectId:   projectID,
		PipelineId:  pipelineID,
		ExecutionId: executionID,
	}
	u.fireEventUsecase.PublishUploadLogsEvent(ctx, &message)
	log.Info(ctx, "FireEventUploadFileLogs success")
	return nil
}

func NewUploadLogWebhookUseCase(fireEventUsecase IFireEventUsecase) IUploadLogWebhookUseCase {
	return &UploadLogWebhookUseCase{
		fireEventUsecase: fireEventUsecase,
	}
}
