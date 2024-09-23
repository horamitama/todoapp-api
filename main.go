package main

import (
	"todoapp-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	r.POST("/signout", controller.SignOut)
	r.POST("/task", controller.CreateTask)
	r.GET("/task/:userId", controller.GetTasks)
	r.Run()
}
