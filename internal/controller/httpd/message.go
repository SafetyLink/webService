package httpd

import (
	"github.com/SafetyLink/commons/jwt"
	"github.com/gofiber/fiber/v2"
)

func (s *WebServiceHttpServer) MountMessageRoute(app *fiber.App) {
	chat := app.Group("/v1/chat")

	chat.Use(jwt.New(jwt.Config{}))

	chat.Post("/:chatID/message", s.createMessage)
}
