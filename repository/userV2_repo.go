package repository

import (
	"context"
	"helloword/model"

	"gorm.io/gorm"
)

type UserV2Repository interface {
	GetUser(
		ctx context.Context,
		id int,
	) (model.User, error)
	GetAllutilisateur(
		ctx context.Context,
	) ([]model.User, error)
}
type userV2Repository struct {
	DB *gorm.DB
}

// GetAllutilisateur implements [UserV2Repository].
func (u *userV2Repository) GetAllutilisateur(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := u.DB.WithContext(ctx).
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser implements [UserV2Repository].
func (u *userV2Repository) GetUser(ctx context.Context, id int) (model.User, error) {
	var user model.User
	err := u.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// constructeur
func NewUserV2Repository(repo *gorm.DB) UserV2Repository {
	return &userV2Repository{
		DB: repo,
	}
}
