package httpd

import "github.com/gofiber/fiber/v2"

func (s *WebServiceHttpServer) MountAuthRoute(app *fiber.App) {
	auth := app.Group("/v1/auth")

	auth.Post("/login", s.login)

}
