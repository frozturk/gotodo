package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
	td, err := auth.CreateToken(user.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = auth.SaveToken(user.ID, td)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func Logout(c *gin.Context) {
	ad, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	deleted, err := auth.DeleteToken(ad.AccessUuid)
	if err != nil || deleted == 0 {
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
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	refreshToken := mapToken["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("RTSECRET")), nil
	})

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && token.Valid {
		c.Status(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		deleted, err := auth.DeleteToken(refreshUuid)
		if err != nil || deleted == 0 {
			c.Status(http.StatusForbidden)
			return
		}

		ts, err := auth.CreateToken(userId)
		if err != nil {
			c.Status(http.StatusForbidden)
			return
		}

		err = auth.SaveToken(userId, ts)
		if err != nil {
			c.Status(http.StatusForbidden)
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusOK, tokens)
	} else {
		c.Status(http.StatusUnauthorized)
	}

}

func getUser(c *gin.Context) (*models.User, error) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
