package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib-security/web/config"
	"github.com/golibs-starter/golib/web/actuator"
	"github.com/khaitq-vnist/auto_ci_be/public/controller"
	"github.com/khaitq-vnist/auto_ci_be/public/middleware"
	"go.uber.org/fx"
	"time"
)

type RegisterRoutersIn struct {
	fx.In
	App                   *golib.App
	Engine                *gin.Engine
	Actuator              *actuator.Endpoint
	IntegrationController *controller.IntegrationController
	RepositoryController  *controller.RepositoryController
	ProjectController     *controller.ProjectController
	PipelineController    *controller.PipelineController
	WebHookController     *controller.WebHookController
	UserController        *controller.UserController
	SecurityProperties    *config.HttpSecurityProperties
	ServiceController     *controller.ServiceController
}

func RegisterGinRouters(p RegisterRoutersIn) {
	p.Engine.Use(
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
			AllowHeaders: []string{
				"Origin",
				"Content-Type",
				"Accept",
				"User-Agent",
				"Referer",
				"sec-ch-ua",
				"sec-ch-ua-platform",
				"sec-ch-ua-mobile",
			},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))
	v1UserGroup := group.Group("/v1/auth")
	{
		v1UserGroup.POST("/sign-up", p.UserController.CreateUser)
		v1UserGroup.POST("/login", p.UserController.LoginUser)
	}
	v1IntegrationGroup := group.Group("/v1/integrations", middleware.GetInfoFromToken(p.SecurityProperties.Jwt))
	{
		v1IntegrationGroup.POST("", p.IntegrationController.CreateIntegration)
		v1IntegrationGroup.GET("", p.IntegrationController.GetIntegration)
	}
	v1RepositoryGroup := group.Group("/v1/repositories")
	{
		v1RepositoryGroup.GET("/integration/:integrationId", p.RepositoryController.GetRepositoriesByIntegrationId)
	}
	v1ProjectGroup := group.Group("/v1/projects")
	{
		v1ProjectGroup.GET("", p.ProjectController.GetProjectList)
		v1ProjectGroup.POST("", p.ProjectController.CreateProject)
		v1ProjectGroup.GET("/:projectId/analyze", p.ProjectController.AnalyzeProject)
		v1ProjectGroup.GET("/:projectId/branches", p.ProjectController.GetListBranches)
		v1ProjectGroup.GET("/template/:buildTool", p.ProjectController.GetTemplateByBuildTool)
		v1ProjectGroup.GET("/:projectId/pipelines", p.PipelineController.GetListPipeline)
		v1ProjectGroup.GET("/:projectId/pipelines/:pipelineId/executions", p.PipelineController.GetListExecution)
		v1ProjectGroup.GET("/:projectId/pipelines/:pipelineId/executions/:executionId", p.PipelineController.GetExecutionDetail)
		v1ProjectGroup.POST("/:projectId/pipelines/:pipelineId/executions", p.PipelineController.RunExecution)
		v1ProjectGroup.DELETE("/:projectId/pipelines/:pipelineId", p.PipelineController.DeletePipeline)
	}
	v1PipelineGroup := group.Group("/v1/pipelines")
	{
		v1PipelineGroup.POST("", p.PipelineController.CreatePipeline)
	}
	v1WebHookGroup := group.Group("/v1/webhooks")
	{
		v1WebHookGroup.GET("/upload-logs/projects/:projectId/pipelines/:pipelineId/executions/:executionId", p.WebHookController.HandleUploadFileLogs)
	}
	v1ServiceGroup := group.Group("/v1/services")
	{
		v1ServiceGroup.GET("", p.ServiceController.GetAllServices)
	}
}
