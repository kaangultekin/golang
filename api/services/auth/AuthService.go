package auth

import (
	userInterfaces "golang/api/interfaces/user"
	authFormStructs "golang/api/structs/form/auth"
	resultStructs "golang/api/structs/result"
)

type AuthService struct {
	UserRepository userInterfaces.IUserRepository
}

func (as *AuthService) GetUserByID(id int) (*resultStructs.ResultStruct, error) {
	user, err := as.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (as *AuthService) Register(registerForm authFormStructs.RegisterFormStruct) (interface{}, error) {
	return registerForm, nil
}
