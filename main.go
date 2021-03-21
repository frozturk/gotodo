package main

import (
	"github.com/frozturk/gologin/controllers"
	"github.com/frozturk/gologin/db"
	"github.com/frozturk/gologin/middleware"
	_ "github.com/frozturk/gologin/redis"
	"github.com/gin-gonic/gin"
)

func main() {
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
