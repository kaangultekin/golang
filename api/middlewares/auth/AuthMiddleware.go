package auth

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	messageConstants "golang/api/constants/message"
	resultStructs "golang/api/structs/result"
	"os"
)

func AuthMiddleware() fiber.Handler {
	var (
		result *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)
	jwtSecret := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			result.Success = false
			result.Code = fiber.StatusUnauthorized
			result.Message = messageConstants.ErrInvalidToken

			return c.Status(result.Code).JSON(result)
		},
		SigningKey: []byte(jwtSecret),
	})
}
