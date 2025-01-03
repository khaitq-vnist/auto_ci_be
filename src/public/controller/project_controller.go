package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/middleware"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/request"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/response"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
	"strconv"
)

type ProjectController struct {
	projectService service.IProjectService
}

func NewProjectController(projectService service.IProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}
func (p ProjectController) GetProjectList(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "get user id error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	result, err := p.projectService.GetProjectList(c, userID)
	if err != nil {
		log.Error(c, "get project list error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToListProjectResponse(result))
}
func (p ProjectController) CreateProject(c *gin.Context) {
	userId, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "get user id error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	var request request.CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error(c, "bind request error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := p.projectService.CreateProject(c, userId, request.IntegrationId, request.RepoId)
	if err != nil {
		log.Error(c, "create project error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToProjectResponse(result))
}
func (p ProjectController) AnalyzeProject(c *gin.Context) {
	userId := int64(1)
	projectId, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse project id error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := p.projectService.AnalyzeProject(c, userId, projectId)
	if err != nil {
		log.Error(c, "analyze project error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToProjectAnalyzeResponse(result.Tools))
}
func (p ProjectController) GetListBranches(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "get user id error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	projectId, err := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err != nil {
		log.Error(c, "parse project id error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	result, err := p.projectService.GetListBranches(c, userID, projectId)
	if err != nil {
		log.Error(c, "get list branches error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToListBranchResponse(result))
}

func (p ProjectController) GetTemplateByBuildTool(c *gin.Context) {
	buildTool := c.Param("buildTool")
	result, err := p.projectService.GetTemplateByBuildTool(c, buildTool)
	if err != nil {
		log.Error(c, "get template by build tool error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)
}
