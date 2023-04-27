package user

import (
	"golang/api/config"
	UserModel "golang/api/models/user"
	"golang/api/structs/result"
)

type UserRepository struct{}

func (ur *UserRepository) GetByID(id int) (*result.Result, error) {
	result := &result.Result{
		Datas: &UserModel.User{},
	}

	config.DB.Table("users").
		Select("id, status, name, surname, email, created_at, updated_at").
		Where("id = ?", id).
		Order("id desc").
		First(&result.Datas)

	return result, nil
}
