package bootstrap

import (
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
	"github.com/khaitq-vnist/auto_ci_be/worker/handler"
	"github.com/khaitq-vnist/auto_ci_be/worker/router"
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
		golib.HttpClientOpt(),
		golibsec.SecuredHttpClientOpt(),

		golibdata.RedisOpt(),
		golibdata.DatasourceOpt(),

		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaAdminOpt(),
		golibmsg.KafkaConsumerOpt(),

		//Provide properties
		//golib.ProvideProps(properties.NewGitlabProperties),
		golib.ProvideProps(properties.NewGithubProperties),
		golib.ProvideProps(properties2.NewEncryptProperties),
		golib.ProvideProps(properties.NewBuddyProperties),
		golib.ProvideProps(properties.NewGCSProperties),

		//Provide port's implements
		fx.Provide(postgres.NewBaseRepository),
		fx.Provide(postgres.NewUserRepositoryAdapter),
		fx.Provide(postgres.NewProviderRepositoryAdapter),
		fx.Provide(postgres.NewIntegrationRepositoryAdapter),
		fx.Provide(client.NewThirdPartyProviderAdapter),
		fx.Provide(postgres.NewProjectRepositoryAdapter),
		fx.Provide(publisher.NewEventPublisherAdapter),

		fx.Provide(client.NewGithubProviderClient),
		fx.Provide(client.NewThirdPartyToolAdapter),
		fx.Provide(client.NewGCSClientAdapter),

		//Provide usecase
		fx.Provide(usecase.NewGetThirdPartyProviderUseCase),
		fx.Provide(usecase.NewCreateIntegrationUseCase),
		fx.Provide(usecase.NewEncryptUseCase),
		fx.Provide(usecase.NewGetProviderUseCase),
		fx.Provide(usecase.NewGetIntegrationUseCase),
		fx.Provide(usecase.NewGetRepositoryUseCase),
		fx.Provide(usecase.NewGetProjectUseCase),
		fx.Provide(usecase.NewUploadLogWebhookUseCase),
		fx.Provide(usecase.NewFireEventUsecase),
		fx.Provide(usecase.NewUploadFileLogUseCase),

		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),
		//provide Handlers
		golibmsg.ProvideConsumer(handler.NewUploadFileLogsHandler),

		golibgin.OnStopHttpServerOpt(),

		golibmsg.OnStopConsumerOpt(),
	)
}
