package setup

import (
	"helloword/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	handler.ParamUserRoute(r, db)
	return r
}
