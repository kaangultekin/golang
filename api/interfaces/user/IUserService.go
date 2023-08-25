package user

import (
	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	GetUser(id int) (interface{}, error)
	SearchUsers(query string, c *fiber.Ctx) (interface{}, error)
}
