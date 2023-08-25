package user

import (
	"github.com/gofiber/fiber/v2"
	messageConstants "golang/api/constants/message"
	userInterfaces "golang/api/interfaces/user"
	resultStructs "golang/api/structs/result"
	"strconv"
)

type UserController struct {
	UserService userInterfaces.IUserService
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	var (
		result *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
	)

	id, _ := strconv.Atoi(c.Params("id"))

	getUser, err := uc.UserService.GetUser(id)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusBadRequest
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusOK
	result.Message = messageConstants.SuccessGeneralMessage
	result.Datas = getUser

	return c.Status(result.Code).JSON(result)
}

func (uc *UserController) SearchUsers(c *fiber.Ctx) error {
	var (
		result *resultStructs.ResultStruct = &resultStructs.ResultStruct{}
		query                              = c.Params("query")
	)

	searchUsers, err := uc.UserService.SearchUsers(query, c)

	if err != nil {
		result.Success = false
		result.Code = fiber.StatusBadRequest
		result.Message = err.Error()

		return c.Status(result.Code).JSON(result)
	}

	result.Success = true
	result.Code = fiber.StatusOK
	result.Message = messageConstants.SuccessGeneralMessage
	result.Datas = searchUsers

	return c.Status(result.Code).JSON(result)
}
