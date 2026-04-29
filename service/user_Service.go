package service

import (
	"helloword/model"
	"helloword/repository"
)

type UserService interface {
	Save(user model.User) error
	GetAllUser() ([]model.User, error)
}
type userService struct {
	service repository.UserRepository
}

// GetAllUser implements [UserService].
func (u *userService) GetAllUser() ([]model.User, error) {
	return u.service.GetAllUser()
}

// Save implements [UserService].
func (u *userService) Save(user model.User) error {
	return u.service.Save(user)
}

//on creer le constructeur

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		service: repo,
	}
}
