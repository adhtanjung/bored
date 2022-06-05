package handlers

import (
	"log"
	"net/http"
	"net/mail"

	database "github.com/adhtanjung/bored/pkg/common/db"
	"github.com/adhtanjung/bored/pkg/common/models"
	"github.com/adhtanjung/bored/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

func emailValidator(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil

}

type LoginInput struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {

	db := database.DB.Db
	var user models.User

	var input LoginInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var query string
	if emailValidator(input.Identity) {
		query = "email= ?"
	} else {
		query = "username= ?"
	}

	if err := db.Where(query, input.Identity).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	pass := input.Password
	log.Println("pass:", pass)
	log.Println("userPass:", user.Password)

	if !utils.ValidatePassword(pass, user.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error", "message": "Password incorrect",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status": "success", "message": "Success login", "token": token,
	})
}
