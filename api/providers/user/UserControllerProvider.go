package providers

import (
	userControllers "golang/api/controllers/user"
	userInterfaces "golang/api/interfaces/user"
)

func UserControllerProvider(userService userInterfaces.IUserService) *userControllers.UserController {
	return &userControllers.UserController{UserService: userService}
}
