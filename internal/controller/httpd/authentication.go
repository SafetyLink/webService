package httpd

import "github.com/gofiber/fiber/v2"

func (s *WebServiceHttpServer) MountAuthRoute(app *fiber.App) {
	chat := app.Group("/v1/auth")

	chat.Post("/login")
}
