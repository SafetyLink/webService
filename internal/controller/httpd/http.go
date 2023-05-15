package httpd

import (
	"github.com/SafetyLink/commons/otel"
	"github.com/SafetyLink/webService/internal/domain/message"
	"github.com/SafetyLink/webService/internal/domain/repo"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type WebServiceHttpServer struct {
	MessageSrv   message.Srv
	GrpcUserRepo repo.User
	GrpcAuthRepo repo.Authentication
}

func NewWebServiceHttpServer(MessageSrv message.Srv, GrpcUserRepo repo.User, GrpcAuthRepo repo.Authentication) *WebServiceHttpServer {
	return &WebServiceHttpServer{
		MessageSrv:   MessageSrv,
		GrpcUserRepo: GrpcUserRepo,
		GrpcAuthRepo: GrpcAuthRepo,
	}
}

func (s *WebServiceHttpServer) Router() *fiber.App {
	_ = otel.InitHttpTracer()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(otelfiber.Middleware())
	s.MountMessageRoute(app)

	s.MountAuthRoute(app)

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
