package controller

import (
	"helloword/model"
	"helloword/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Save(ctx *gin.Context) error
	GetAllUser() ([]model.User, error)
}

type userController struct {
	controller service.UserService
}

// GetAllUser implements [UserController].
func (u *userController) GetAllUser() ([]model.User, error) {
	return u.controller.GetAllUser()
}

// Save implements [UserController].
func (u *userController) Save(ctx *gin.Context) error {
	var user model.User
	ctx.ShouldBindJSON(&user)
	return u.controller.Save(user)
}

//constructeur

func NewUserController(service service.UserService) UserController {
	return &userController{
		controller: service,
	}
}
