package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib/web/actuator"
	"github.com/khaitq-vnist/auto_ci_be/public/controller"
	"go.uber.org/fx"
	"time"
)

type RegisterRoutersIn struct {
	fx.In
	App                   *golib.App
	Engine                *gin.Engine
	Actuator              *actuator.Endpoint
	IntegrationController *controller.IntegrationController
}

func RegisterGinRouters(p RegisterRoutersIn) {
	p.Engine.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		}))
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))

	v1IntegrationGroup := group.Group("/v1/integrations")
	{
		v1IntegrationGroup.POST("", p.IntegrationController.CreateIntegration)
		v1IntegrationGroup.GET("", p.IntegrationController.GetIntegration)
	}

}
