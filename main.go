package main

import (
	"todoapp-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/task", controller.CreateTask)
	r.GET("/tasks", controller.GetTasks)
	r.Run()
}
