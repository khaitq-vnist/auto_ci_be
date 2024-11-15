package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/resource/response"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
	"strconv"
)

type RepositoryController struct {
	repoService service.IRepositoryService
}

func NewRepositoryController(repoService service.IRepositoryService) *RepositoryController {
	return &RepositoryController{repoService: repoService}
}
func (r RepositoryController) GetRepositoriesByIntegrationId(c *gin.Context) {
	integrationId, err := strconv.ParseInt(c.Param("integrationId"), 10, 64)
	if err != nil {
		log.Error(c, "parse integrationId error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}

	userId := int64(1)
	result, err := r.repoService.GetRepositoriesByIntegrationId(c, integrationId, userId)
	if err != nil {
		log.Error(c, "get repositories error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToListReposResponse(result))
}
