package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/public/apihelper"
	"github.com/khaitq-vnist/auto_ci_be/public/middleware"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
)

type DashboardController struct {
	dashboardService service.IDashboardService
}

func (d DashboardController) GetDashboard(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "get user id error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	result, err := d.dashboardService.GetDashboard(c, userID)
	if err != nil {
		log.Error(c, "get dashboard error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)
}
func NewDashboardController(dashboardService service.IDashboardService) *DashboardController {
	return &DashboardController{
		dashboardService: dashboardService,
	}
}
