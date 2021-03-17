package main

import (
	"github.com/frozturk/gologin/controllers"
	"github.com/frozturk/gologin/db"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(db.Inject())

	app.GET("/todo", controllers.GetAll)
	app.POST("/todo", controllers.Create)

	app.POST("/login", controllers.Login)
	app.POST("/signup", controllers.SignUp)

	app.Run(":4000")
}
