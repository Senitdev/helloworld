package handler

import (
	"helloword/controller"
	"helloword/model"
	"helloword/repository"
	"helloword/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamUserRoute(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	g := r.Group("/api/v2/")
	g.GET("/user", func(ctx *gin.Context) {
		var user []model.User
		user, err := userController.GetAllUser()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(200, user)
	})
	g.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})
}
