package todo

import (
	"net/http"

	"github.com/frozturk/gologin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todos []models.Todo
	if err := db.Find(&todos).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, todos)
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.String(http.StatusBadRequest, "asdas")
		return
	}
	if err = db.Create(&todo).Error; err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, todo)
}
