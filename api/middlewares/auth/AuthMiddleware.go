package auth

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"golang/api/config"
	messageConstants "golang/api/constants/message"
	"golang/api/helpers"
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
		SuccessHandler: func(c *fiber.Ctx) error {
			token := helpers.GetToken(c)

			tokenControl, _ := config.Redis.Get(c.Context(), token).Result()

			if tokenControl != "" {
				result.Success = false
				result.Code = fiber.StatusUnauthorized
				result.Message = messageConstants.ErrInvalidToken

				return c.Status(result.Code).JSON(result)
			}

			return c.Next()
		},
		SigningKey: []byte(jwtSecret),
	})
}
