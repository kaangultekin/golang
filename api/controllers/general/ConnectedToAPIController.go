package general

import (
	"github.com/gofiber/fiber/v2"
	messageConstants "golang/api/constants/message"
	resultStructs "golang/api/structs/result"
)

type ConnectedToAPIController struct{}

func (cac *ConnectedToAPIController) ConnectedToAPI(c *fiber.Ctx) error {
	result := &resultStructs.ResultStruct{
		Success: true,
		Code:    fiber.StatusOK,
		Message: messageConstants.SuccessAPI,
	}

	return c.JSON(result)
}
