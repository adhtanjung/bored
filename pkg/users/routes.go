package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/adhtanjung/bored/pkg/users/handlers"
)

type handler struct {
	DB *gorm.DB
}

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")
	api.POST("/login", handlers.Login)
	user := api.Group("/users")
	user.POST("/", handlers.AddUser)
	user.GET("/", handlers.GetUsers)
}
