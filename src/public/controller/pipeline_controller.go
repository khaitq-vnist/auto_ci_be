package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/request"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/response"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
	"strconv"
)

type PipelineController struct {
	pipelineService service.IPipelineService
}

func (p PipelineController) CreatePipeline(c *gin.Context) {
	projectID, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	var req request.CreatePipelineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorc(c, "bind request error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}

	result, err := p.pipelineService.CreateNewPipeline(c, projectID, request.ToPipelineEntity(&req))
	if err != nil {
		log.Errorc(c, "create pipeline error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)
}
func (p PipelineController) GetListPipeline(c *gin.Context) {
	projectID, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := p.pipelineService.GetListPipelineByProjectID(c, projectID)
	if err != nil {
		log.Error(c, "get list pipeline error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToListPipelineResponse(result))
}
func (p PipelineController) GetListExecution(c *gin.Context) {
	projectId, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	pipelineId, err := strconv.ParseInt(c.Param("pipelineId"), 10, 64)
	if err != nil {
		log.Error(c, "parse pipelineID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := p.pipelineService.GetListExecutions(c, projectId, pipelineId)
	if err != nil {
		log.Error(c, "get list execution error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)
}
func (p PipelineController) GetExecutionDetail(c *gin.Context) {
	projectId, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	pipelineId, err := strconv.ParseInt(c.Param("pipelineId"), 10, 64)
	if err != nil {
		log.Error(c, "parse pipelineID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	executionId, err := strconv.ParseInt(c.Param("executionId"), 10, 64)
	if err != nil {
		log.Error(c, "parse executionID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := p.pipelineService.GetExecutionDetailByID(c, projectId, pipelineId, executionId)
	if err != nil {
		log.Error(c, "get execution detail error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)
}
func (p PipelineController) RunExecution(c *gin.Context) {
	projectId, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	pipelineId, err := strconv.ParseInt(c.Param("pipelineId"), 10, 64)
	if err != nil {
		log.Error(c, "parse pipelineID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := p.pipelineService.RunExecution(c, projectId, pipelineId)
	if err != nil {
		log.Error(c, "run execution error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)
}
func (p PipelineController) DeletePipeline(c *gin.Context) {
	projectId, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	pipelineId, err := strconv.ParseInt(c.Param("pipelineId"), 10, 64)
	if err != nil {
		log.Error(c, "parse pipelineID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	err = p.pipelineService.DeletePipelineByID(c, projectId, pipelineId)
	if err != nil {
		log.Error(c, "delete pipeline error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
func (p PipelineController) GetDetailLog(c *gin.Context) {
	projectID, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse projectID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	pipelineID, err := strconv.ParseInt(c.Param("pipelineId"), 10, 64)
	if err != nil {
		log.Error(c, "parse pipelineID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	executionID, err := strconv.ParseInt(c.Param("executionId"), 10, 64)
	if err != nil {
		log.Error(c, "parse executionID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	actionID, err := strconv.ParseInt(c.Param("actionId"), 10, 64)
	if err != nil {
		log.Error(c, "parse actionID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	rsp, err := p.pipelineService.GetDetailLog(c, projectID, pipelineID, executionID, actionID)
	if err != nil {
		log.Error(c, "get detail log error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, rsp)
}
func NewPipelineController(pipelineService service.IPipelineService) *PipelineController {
	return &PipelineController{
		pipelineService: pipelineService,
	}
}
