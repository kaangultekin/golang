package user

import (
	UserModel "golang/api/models/user"
)

type IUserRepository interface {
	GetByID(id int) (*UserModel.User, error)
}
