package httpd

import (
	"github.com/SafetyLink/commons/api/response"
	"github.com/gofiber/fiber/v2"
)

func (s *WebServiceHttpServer) getUserProfileByID(c *fiber.Ctx) error {
	profileID, err := c.ParamsInt("profileID")
	if err != nil {
		return err
	}

	user, err := s.GrpcUserRepo.GetUserByID(c.Context(), int64(profileID))
	if err != nil {
		return err
	}

	response.SuccessDataJson(c, 200, user)
	return nil
}

func (s *WebServiceHttpServer) getSelf(c *fiber.Ctx) error {
	profile, err := s.GrpcUserRepo.GetSelf(c.Context())
	if err != nil {
		return err
	}
	profile.Security = nil

	response.SuccessDataJson(c, 200, profile)
	return nil
}
