package httpd

import (
	"github.com/SafetyLink/commons/jwt"
	"github.com/gofiber/fiber/v2"
)

func (s *WebServiceHttpServer) MountProfileRoutes(app *fiber.App) {
	profile := app.Group("/v1/profile")

	profile.Use(jwt.New(jwt.Config{}))

	profile.Get("/self", s.getSelf)
	profile.Get("/:profileID", s.getUserProfileByID)

}
