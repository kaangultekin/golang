package auth

import (
	userInterfaces "golang/api/interfaces/user"
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
