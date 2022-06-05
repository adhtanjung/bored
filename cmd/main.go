package main

import (
	database "github.com/adhtanjung/bored/pkg/common/db"
	"github.com/adhtanjung/bored/pkg/users"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port": "asd",
		})
	})

	users.SetupRoutes(r)

	r.Run(":8080")
}
