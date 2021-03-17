package controllers

import (
	"net/http"

	"github.com/frozturk/gologin/auth"
	"github.com/frozturk/gologin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := db.Where(&models.User{Username: user.Username, Password: user.Password}).First(&user).Error; err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, token)
}

func SignUp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user, err := getUser(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func getUser(c *gin.Context) (*models.User, error) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
