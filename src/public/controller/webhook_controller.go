package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
	"strconv"
)

type WebHookController struct {
	webhookService service.IWebhookService
}

func NewWebHookController(webhookService service.IWebhookService) *WebHookController {
	return &WebHookController{webhookService: webhookService}
}
func (w WebHookController) HandleUploadFileLogs(c *gin.Context) {
	projectId, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectId error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	pipelineId, err := strconv.ParseInt(c.Param("pipelineId"), 10, 64)
	if err != nil {
		log.Error(c, "parse pipelineId error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	executionId, err := strconv.ParseInt(c.Param("executionId"), 10, 64)
	if err != nil {
		log.Error(c, "parse executionId error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	err = w.webhookService.HandleUploadFileLogs(c, projectId, pipelineId, executionId)
	if err != nil {
		log.Error(c, "handle upload file logs error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
