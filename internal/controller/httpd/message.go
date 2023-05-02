package httpd

import "github.com/gofiber/fiber/v2"

func (s *WebServiceHttpServer) MountMessageRoute(app *fiber.App) {
	chat := app.Group("/v1/chat")

	chat.Post("/:chatID/message", s.createMessage)
}
