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

type IntegrationController struct {
	integrationService service.IIntegrationService
}

func NewIntegrationController(integrationService service.IIntegrationService) *IntegrationController {
	return &IntegrationController{
		integrationService: integrationService,
	}
}

func (i IntegrationController) CreateIntegration(c *gin.Context) {
	var request request.CreateIntegrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error(c, "bind request error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}

	err := i.integrationService.CreateIntegration(c, &request)
	if err != nil {
		log.Error(c, "create integration error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
func (i IntegrationController) GetIntegration(c *gin.Context) {
	userId := int64(1)
	integrations, err := i.integrationService.GetIntegrationByUserId(c, userId)
	if err != nil {
		log.Error(c, "get integration error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToListIntegrationResponse(integrations))
}
