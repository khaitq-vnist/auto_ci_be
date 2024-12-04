package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/request"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
)

type PipelineController struct {
	pipelineService service.IPipelineService
}

func (p PipelineController) CreatePipeline(c *gin.Context) {
	var req request.CreatePipelineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorc(c, "bind request error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
	}

	result, err := p.pipelineService.CreateNewPipeline(c, 1, request.ToPipelineEntity(&req))
	if err != nil {
		log.Errorc(c, "create pipeline error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
	}
	apihelper.SuccessfulHandle(c, result)
}
func NewPipelineController(pipelineService service.IPipelineService) *PipelineController {
	return &PipelineController{
		pipelineService: pipelineService,
	}
}
