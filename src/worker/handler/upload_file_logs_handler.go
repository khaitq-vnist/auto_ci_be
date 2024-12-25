package handler

import (
	"context"
	"encoding/json"
	"github.com/golibs-starter/golib-message-bus/kafka/core"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/constant"
	"github.com/khaitq-vnist/auto_ci_be/core/event"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
)

type UploadFileLogsHandler struct {
	uploadFileLogUseCase usecase.IUploadFileLogUseCase
}

func (u UploadFileLogsHandler) HandlerFunc(message *core.ConsumerMessage) {
	var e event.UploadFileLogsEvent
	if err := json.Unmarshal(message.Value, &e); err != nil {
		log.Error("[UploadFileLogsHandler] Error when unmarshal event message, detail: ", err)
		return
	}
	if e.AbstractEvent == nil || e.AbstractEvent.ApplicationEvent == nil || e.AbstractEvent.Event != constant.UploadFileLogsEvent || e.PayloadData == nil {
		if message == nil {
			log.Info("[UploadFileLogsHandler] message is nil")
			return
		}
		log.Info("skip key  %s", string(message.Key))
		return
	}
	ctx := context.Background()
	msg := e.PayloadData
	err := u.uploadFileLogUseCase.UploadFileLogByExecutionID(ctx, msg.ProjectId, msg.PipelineId, msg.ExecutionId)
	if err != nil {
		log.Error(ctx, "[UploadFileLogsHandler] Error when upload file logs [%s] to cloud storage [%v]",
			msg.ExecutionId, err)
		return
	}
	log.Info(ctx, "[UploadFileLogsHandler] Success to upload file logs [%s] to cloud storage", e.PayloadData.ExecutionId)

}

func (u UploadFileLogsHandler) Close() {

}

func NewUploadFileLogsHandler(uploadFileLogUseCase usecase.IUploadFileLogUseCase) core.ConsumerHandler {
	return &UploadFileLogsHandler{uploadFileLogUseCase: uploadFileLogUseCase}
}
