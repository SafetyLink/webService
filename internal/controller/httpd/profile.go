package httpd

import "github.com/gofiber/fiber/v2"

func (s *WebServiceHttpServer) MountProfileRoutes(app *fiber.App) {
	profile := app.Group("/v1/profile")

	profile.Get("/:profileID", s.getProfile)
}
