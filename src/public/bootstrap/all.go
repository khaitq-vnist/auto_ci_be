package bootstrap

import (
	"github.com/go-resty/resty/v2"
	"github.com/golibs-starter/golib"
	golibdata "github.com/golibs-starter/golib-data"
	golibgin "github.com/golibs-starter/golib-gin"
	golibmsg "github.com/golibs-starter/golib-message-bus"
	golibsec "github.com/golibs-starter/golib-security"
	"github.com/khaitq-vnist/auto_ci_be/adapter/http/client"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/adapter/publisher"
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
		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaAdminOpt(),
		golibmsg.KafkaProducerOpt(),

		// Provide http client auto config with contextual http client by default,
		// Besides, provide an additional wrapper to easy to control security.
		golib.HttpClientOpt(),
		golibsec.SecuredHttpClientOpt(),

		//Provide properties
		//golib.ProvideProps(properties.NewGitlabProperties),
		golib.ProvideProps(properties.NewGithubProperties),
		golib.ProvideProps(properties2.NewEncryptProperties),
		golib.ProvideProps(properties.NewBuddyProperties),
		golib.ProvideProps(properties.NewSonarProperties),

		//Provide core properties
		golib.ProvideProps(properties2.NewTokenProperties),
		//Provide port's implements
		fx.Provide(postgres.NewBaseRepository),
		fx.Provide(postgres.NewUserRepositoryAdapter),
		fx.Provide(postgres.NewProviderRepositoryAdapter),
		fx.Provide(postgres.NewIntegrationRepositoryAdapter),
		fx.Provide(client.NewThirdPartyProviderAdapter),
		fx.Provide(postgres.NewProjectRepositoryAdapter),
		fx.Provide(postgres.NewPipelineTemplateRepositoryAdapter),
		fx.Provide(postgres.NewPipelineStageTemplateRepoAdapter),
		fx.Provide(postgres.NewStageTemplateRepoAdapter),
		fx.Provide(postgres.NewCommandTemplateRepoAdapter),
		fx.Provide(postgres.NewVariableTemplateRepoAdapter),
		fx.Provide(publisher.NewEventPublisherAdapter),
		fx.Provide(postgres.NewDatabaseTransactionAdapter),
		fx.Provide(postgres.NewServiceRepositoryAdapter),
		fx.Provide(client.NewSonarClientAdapter),
		fx.Provide(postgres.NewExecutionHistoryRepositoryAdapter),

		//Provide client's implements
		fx.Provide(client.NewGithubProviderClient),
		fx.Provide(client.NewThirdPartyToolAdapter),

		//Provide usecase
		fx.Provide(usecase.NewGetThirdPartyProviderUseCase),
		fx.Provide(usecase.NewCreateIntegrationUseCase),
		fx.Provide(usecase.NewEncryptUseCase),
		fx.Provide(usecase.NewGetProviderUseCase),
		fx.Provide(usecase.NewGetIntegrationUseCase),
		fx.Provide(usecase.NewGetRepositoryUseCase),
		fx.Provide(usecase.NewGetProjectUseCase),
		fx.Provide(usecase.NewCreateProjectUseCase),
		fx.Provide(usecase.NewAnalyzeSourceCodeUsecase),
		fx.Provide(usecase.NewGetBranchUseCase),
		fx.Provide(usecase.NewGetPipelineTemplateUsecase),
		fx.Provide(usecase.NewCreatePipelineUsecase),
		fx.Provide(usecase.NewGetPipelineUseCase),
		fx.Provide(usecase.NewGetExecutionUsecase),
		fx.Provide(usecase.NewRunExecutionUsecase),
		fx.Provide(usecase.NewUploadLogWebhookUseCase),
		fx.Provide(usecase.NewFireEventUsecase),
		fx.Provide(usecase.NewDeletePipelineUsecase),
		fx.Provide(usecase.NewCreateUserUseCase),
		fx.Provide(usecase.NewGetUserUseCase),
		fx.Provide(usecase.NewDatabaseTransactionUsecase),
		fx.Provide(usecase.NewLoginUseCase),
		fx.Provide(usecase.NewGetServiceUseCase),
		fx.Provide(usecase.NewGetDetailLogUseCase),
		fx.Provide(usecase.NewGetDashboardUsecase),

		//Provide service
		fx.Provide(service.NewIntegrationService),
		fx.Provide(service.NewRepositoryService),
		fx.Provide(service.NewProjectService),
		fx.Provide(service.NewPipelineService),
		fx.Provide(service.NewWebhookService),
		fx.Provide(service.NewUserService),
		fx.Provide(service.NewServiceService),
		fx.Provide(service.NewDashboardService),

		//Provide controller
		fx.Provide(controller.NewIntegrationController),
		fx.Provide(controller.NewRepositoryController),
		fx.Provide(controller.NewProjectController),
		fx.Provide(controller.NewPipelineController),
		fx.Provide(controller.NewWebHookController),
		fx.Provide(controller.NewUserController),
		fx.Provide(
			func() *resty.Client {
				return resty.New()
			},
		),
		fx.Provide(controller.NewServiceController),
		fx.Provide(controller.NewDashboardController),
		// Provide gin http server auto config,
		// actuator endpoints and application routers
		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),
	)
}
