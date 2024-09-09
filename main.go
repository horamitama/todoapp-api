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
	r.GET("/tasks", controller.GetTasks)
	r.PUT("/task/:id", controller.UpdateTask)
	r.DELETE("/task/:id", controller.DeleteTask)
	r.Run()
}
