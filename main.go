package main

import (
	"fmt"

	todo "github.com/frozturk/gologin/controllers"
	"github.com/frozturk/gologin/db"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("a")
	app := gin.Default()
	app.Use(db.Inject())

	app.GET("/todo", todo.GetAll)
	app.POST("/todo", todo.Create)
	fmt.Println("b")
	app.Run(":4000")
}
