package auth

import userModel "golang/models/user"

type IAuthService interface {
	GetUserByID(id int) (*userModel.User, error)
}
