package httpd

import (
	"github.com/SafetyLink/commons/api/response"
	"github.com/gofiber/fiber/v2"
)

func (s *WebServiceHttpServer) getProfile(c *fiber.Ctx) error {
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
