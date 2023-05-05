package httpd

import (
	"github.com/SafetyLink/webService/internal/domain/message"
	"github.com/SafetyLink/webService/internal/domain/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app := fiber.New()
	app.Use(cors.New())

	s.MountMessageRoute(app)

	s.MountProfileRoutes(app)

	return app
}
