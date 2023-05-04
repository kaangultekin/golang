package general

import (
	"github.com/gofiber/fiber/v2"
	resultStructs "golang/api/structs/result"
)

type ConnectedToAPIController struct{}

func (cac *ConnectedToAPIController) ConnectedToAPI(c *fiber.Ctx) error {
	result := &resultStructs.ResultStruct{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "Connected to API.",
	}

	return c.JSON(result)
}
