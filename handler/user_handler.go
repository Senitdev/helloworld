package handler

import (
	"context"
	"helloword/controller"
	"helloword/model"
	"helloword/repository"
	"helloword/service"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamUserRoute(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewUserV2Repository(db)
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
		err := userController.Save(ctx)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "Ajouter avec succes")
	})
	g.GET("/utilisateur", GetUserHandler(repo))
	g.GET("/utilisateur/:id", GetUserById(repo))
}
func GetUserHandler(repos repository.UserV2Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			2*time.Second,
		)
		defer cancel()
		var user []model.User
		user, err := repos.GetAllutilisateur(ctx)
		if err != nil {
			c.JSON(404, gin.H{
				"error": "Erreur",
			})
			return
		}
		c.JSON(200, user)
	}
}

var sf singleflight.Group

func GetUserById(
	repos repository.UserV2Repository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ids, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Manque le id"})
			return
		}
		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			2*time.Second,
		)
		defer cancel()

		key := "user:" + id

		v, err, shared := sf.Do(
			key,
			func() (interface{}, error) {
				// un seul goroutine arrive ici
				user, err := repos.GetUser(ctx, ids)
				if err != nil {
					return nil, err
				}
				return user, nil
			},
		)
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}
		// optionnel pour debug
		if shared {
			// requête dédupliquée
			//log.Println("singleflight hit")
		}

		user := v.(model.User)

		c.JSON(
			http.StatusOK,
			user,
		)
	}
}
