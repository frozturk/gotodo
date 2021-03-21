package controllers

import (
	"fmt"
	"net/http"

	"github.com/frozturk/gotodo/models"
	"github.com/frozturk/jwtauth"
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
	td, err := jwtauth.CreateToken(user.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, td)
}

func Logout(c *gin.Context) {
	err := jwtauth.LogoutToken(c.Request)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusUnauthorized)
		return
	}
	c.Status(http.StatusOK)
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

func Refresh(c *gin.Context) {
	tokens, err := jwtauth.RefreshTokens(c.Request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, tokens)

}

func getUser(c *gin.Context) (*models.User, error) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
