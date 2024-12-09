package bootstrap

import (
	"github.com/golibs-starter/golib"
	golibdata "github.com/golibs-starter/golib-data"
	golibgin "github.com/golibs-starter/golib-gin"
	golibmsg "github.com/golibs-starter/golib-message-bus"
	golibsec "github.com/golibs-starter/golib-security"
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
		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),
		//provide Handlers
		golibmsg.ProvideConsumer(handler.NewUploadFileLogsHandler),

		golibgin.OnStopHttpServerOpt(),

		golibmsg.OnStopConsumerOpt(),
	)
}
