package db

import (
	"fmt"
	"os"

	"github.com/frozturk/gologin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPASS"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DATABASE CONNECTION FAILURE!")
	}
	db = database

	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.User{})
}

func Inject() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func GetDB() *gorm.DB {
	return db
}
