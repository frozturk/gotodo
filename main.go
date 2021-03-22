package main

import (
	"os"
	"time"

	"github.com/frozturk/gotodo/controllers"
	"github.com/frozturk/gotodo/db"
	"github.com/frozturk/gotodo/middleware"
	"github.com/frozturk/gotodo/redis"
	"github.com/frozturk/jwtauth"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	jwtauth.Setup(os.Getenv("JWTSECRET"), redis.GetRedisClient(), time.Minute*7, time.Hour*24*7)

	app := gin.Default()
	app.Use(db.Inject())

	app.GET("/todo", middleware.AuthenticationMW(), controllers.GetAll)
	app.POST("/todo", controllers.Create)

	app.POST("/login", controllers.Login)
	app.GET("/logout", controllers.Logout)
	app.POST("/signup", controllers.SignUp)
	app.POST("/refresh", controllers.Refresh)

	app.Run(":4000")
}
