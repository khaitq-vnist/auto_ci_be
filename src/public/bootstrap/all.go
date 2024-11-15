package bootstrap

import (
	"github.com/golibs-starter/golib"
	golibdata "github.com/golibs-starter/golib-data"
	golibgin "github.com/golibs-starter/golib-gin"
	golibsec "github.com/golibs-starter/golib-security"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/client"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres"
	properties2 "github.com/khaitq-vnist/auto_ci_be/core/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/usecase"
	"github.com/khaitq-vnist/auto_ci_be/public/controller"
	"github.com/khaitq-vnist/auto_ci_be/public/router"
	"github.com/khaitq-vnist/auto_ci_be/public/service"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golib.EventOpt(),
		golib.BuildInfoOpt(Version, CommitHash, BuildTime),
		golib.ActuatorEndpointOpt(),
		golib.HttpRequestLogOpt(),

		// Http security auto config and authentication filters
		golibsec.HttpSecurityOpt(),
		golibsec.JwtAuthFilterOpt(),
		// Provide datasource auto config
		golibdata.RedisOpt(),
		golibdata.DatasourceOpt(),

		// Provide http client auto config with contextual http client by default,
		// Besides, provide an additional wrapper to easy to control security.
		golib.HttpClientOpt(),
		golibsec.SecuredHttpClientOpt(),

		//Provide properties
		//golib.ProvideProps(properties.NewGitlabProperties),
		golib.ProvideProps(properties.NewGithubProperties),
		golib.ProvideProps(properties2.NewEncryptProperties),

		//Provide port's implements
		fx.Provide(postgres.NewBaseRepository),
		fx.Provide(postgres.NewUserRepositoryAdapter),
		fx.Provide(postgres.NewProviderRepositoryAdapter),
		fx.Provide(postgres.NewIntegrationRepositoryAdapter),
		fx.Provide(client.NewThirdPartyProviderAdapter),

		//Provide client's implements
		fx.Provide(client.NewGithubProviderClient),

		//Provide usecase
		fx.Provide(usecase.NewGetThirdPartyProviderUseCase),
		fx.Provide(usecase.NewCreateIntegrationUseCase),
		fx.Provide(usecase.NewEncryptUseCase),
		fx.Provide(usecase.NewGetProviderUseCase),
		fx.Provide(usecase.NewGetIntegrationUseCase),
		fx.Provide(usecase.NewGetRepositoryUseCase),

		//Provide service
		fx.Provide(service.NewIntegrationService),
		fx.Provide(service.NewRepositoryService),

		//Provide controller
		fx.Provide(controller.NewIntegrationController),
		fx.Provide(controller.NewRepositoryController),

		// Provide gin http server auto config,
		// actuator endpoints and application routers
		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),
	)
}
