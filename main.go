package main

import (
	"context"
	"github.com/SafetyLink/commons/config"
	"github.com/SafetyLink/commons/logger"
	"github.com/SafetyLink/commons/otel"
	"github.com/SafetyLink/webService/internal"
	"github.com/SafetyLink/webService/internal/controller/httpd"
	"github.com/SafetyLink/webService/internal/domain/message"
	"github.com/SafetyLink/webService/internal/infra/adapters/clients"
	"github.com/SafetyLink/webService/internal/infra/adapters/rabbitMQ"
	"github.com/SafetyLink/webService/internal/infra/grpcAuthenticationRepository"
	"github.com/SafetyLink/webService/internal/infra/grpcUserRepository"
	"github.com/SafetyLink/webService/internal/infra/rmqPublisherRepository"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func init() {
	//load env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	fx.New(
		fx.Provide(logger.InitLogger),
		fx.Provide(config.ReadConfig[internal.Config]),
		fx.Provide(otel.InitTracer),

		fx.Provide(rabbitMQ.NewRabbitMQProvider),

		fx.Provide(rmqPublisherRepository.NewRabbitMQPubRepository),

		fx.Provide(message.NewMessageSrv),

		fx.Provide(clients.GrpcAuthenticationClient),

		fx.Provide(grpcUserRepository.NewGrpcUserRepository),
		fx.Provide(grpcAuthenticationRepository.NewGrpcAuthenticationRepository),

		fx.Provide(httpd.NewWebServiceHttpServer),

		fx.Invoke(StartHttpRouter),
	).Run()
}

func StartHttpRouter(lc fx.Lifecycle, httpd *httpd.WebServiceHttpServer, logger *zap.Logger, config *internal.Config) {
	var app *fiber.App
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			app = httpd.Router()
			go func() {
				if err := app.Listen(config.Httpd.Port); err != nil {
					logger.Fatal("Failed to start internal http server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := app.Shutdown()
			if err != nil {
				return err
			}
			return nil
		},
	})
}

func GrpcServerHook() {

}
