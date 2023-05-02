package httpd

import (
	"github.com/SafetyLink/commons/api/response"
	"github.com/SafetyLink/commons/types"
	"github.com/gofiber/fiber/v2"
)

func (s *WebServiceHttpServer) createMessage(c *fiber.Ctx) error {
	var message *types.Message

	if err := c.BodyParser(&message); err != nil {
		response.ErrorJson(c, 400, err.Error())
		return nil
	}

	message, err := s.MessageSrv.CreateMessage(c.Context(), *message)
	if err != nil {
		return err
	}

	response.SuccessDataJson(c, 200, message)
	return nil
}
