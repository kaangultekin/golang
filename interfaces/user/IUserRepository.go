package user

import UserModel "golang/models/user"

type IUserRepository interface {
	GetByID(id int) (*UserModel.User, error)
}
