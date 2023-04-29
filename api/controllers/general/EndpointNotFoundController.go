package general

import (
	"github.com/gofiber/fiber/v2"
	resultStructs "golang/api/structs/result"
)

type EndpointNotFoundController struct{}

func (enfc *EndpointNotFoundController) EndpointNotFound(c *fiber.Ctx) error {
	result := &resultStructs.ResultStruct{
		Success: false,
		Code:    fiber.StatusNotFound,
		Message: "Endpoint not found.",
		Datas:   []string{},
	}

	return c.Status(result.Code).JSON(result)
}
