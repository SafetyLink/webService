package httpd

import "github.com/gofiber/fiber/v2"

func (s *WebServiceHttpServer) MountProfileRoutes(app *fiber.App) {
	chat := app.Group("/v1/profile")

	chat.Post("/:profileID", s.getProfile)
}
