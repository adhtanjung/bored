package handlers

import (
	"net/http"

	database "github.com/adhtanjung/bored/pkg/common/db"
	"github.com/adhtanjung/bored/pkg/common/models"
	"github.com/adhtanjung/bored/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(c *gin.Context) {
	body := AddUserRequestBody{}
	db := database.DB.Db

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := db.Where("username = ?", body.Username).First(&user).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "username already taken", "status": "failed", "data": nil})
		return
	}

	if err := db.Where("email= ?", body.Email).First(&user).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "email already taken", "status": "failed", "data": nil})
		return
	}

	user.Username = body.Username
	user.Email = body.Email
	if hashed, err := utils.HashPassword(body.Password); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		user.Password = hashed
	}

	if result := db.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func GetUsers(c *gin.Context) {
	db := database.DB.Db
	var users []models.User

	// find user in the database
	db.Find(&users)

	if len(users) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Users not found"})
		return
	}

	// else return users

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Success fetched all users", "data": users})

}
