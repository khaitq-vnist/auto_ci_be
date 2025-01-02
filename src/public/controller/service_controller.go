package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/response"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
)

type ServiceController struct {
	serviceService service.IServiceService
}

func (s *ServiceController) GetAllServices(c *gin.Context) {
	result, err := s.serviceService.GetAllService(c)
	if err != nil {
		log.Error(c, "get all services error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromListEntityToServiceResponse(result))
}
func NewServiceController(serviceService service.IServiceService) *ServiceController {
	return &ServiceController{
		serviceService: serviceService,
	}
}
