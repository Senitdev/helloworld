package repository

import (
	"helloword/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) error
	GetAllUser() ([]model.User, error)
}
type userRepostory struct {
	DB *gorm.DB
}

// GetAllUser implements [UserRepository].
func (u *userRepostory) GetAllUser() ([]model.User, error) {
	var user []model.User
	u.DB.Find(&user)
	return user, nil
}

// Save implements [UserRepository].
func (u *userRepostory) Save(user model.User) error {
	if result := u.DB.Create(&user).Error; result != nil {
		return result
	}
	return nil
}

// Constructeur
func NewUserRepository(repo *gorm.DB) UserRepository {
	return &userRepostory{
		DB: repo,
	}
}
