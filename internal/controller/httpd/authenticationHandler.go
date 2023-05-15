package httpd

import (
	"github.com/SafetyLink/commons/api/response"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *WebServiceHttpServer) login(c *fiber.Ctx) error {
	var loginRequest *loginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		response.ErrorJson(c, 400, err.Error())
		return nil
	}

	jwtToken, err := s.GrpcAuthRepo.Login(c.UserContext(), loginRequest.Email, loginRequest.Password)

	if s, ok := status.FromError(err); ok {
		if s.Code() == codes.InvalidArgument {
			response.ErrorJson(c, 400, "invalid email or password")
			return nil
		} else {
			response.ErrorJson(c, 500, err.Error())
			return nil
		}
	}

	response.SuccessDataJson(c, 200, jwtToken)
	return nil
}
