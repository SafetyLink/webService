package httpd

import (
	"github.com/SafetyLink/webService/internal/domain/message"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type WebServiceHttpServer struct {
	MessageSrv message.Srv
}

func NewWebServiceHttpServer(MessageSrv message.Srv) *WebServiceHttpServer {
	return &WebServiceHttpServer{
		MessageSrv: MessageSrv,
	}
}

func (s *WebServiceHttpServer) Router() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	s.MountMessageRoute(app)

	return app
}
