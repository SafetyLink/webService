package httpd

import (
	"context"
	"github.com/SafetyLink/commons/otel"
	"github.com/SafetyLink/webService/internal/domain/message"
	"github.com/SafetyLink/webService/internal/domain/repo"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type WebServiceHttpServer struct {
	MessageSrv   message.Srv
	GrpcUserRepo repo.User
}

func NewWebServiceHttpServer(MessageSrv message.Srv, GrpcUserRepo repo.User) *WebServiceHttpServer {
	return &WebServiceHttpServer{
		MessageSrv:   MessageSrv,
		GrpcUserRepo: GrpcUserRepo,
	}
}

func (s *WebServiceHttpServer) Router() *fiber.App {
	tp := otel.InitTracerHttp()

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(otelfiber.Middleware())

	s.MountMessageRoute(app)

	s.MountProfileRoutes(app)

	app.Get("/health", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("ok")
		if err != nil {
			return err
		}
		return nil
	})

	return app
}
