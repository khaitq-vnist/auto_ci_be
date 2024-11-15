package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/request"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/response"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
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
	userId := int64(1)
	result, err := p.projectService.GetProjectList(c, userId)
	if err != nil {
		log.Error(c, "get project list error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToListProjectResponse(result))
}
func (p ProjectController) CreateProject(c *gin.Context) {
	userId := int64(1)
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
