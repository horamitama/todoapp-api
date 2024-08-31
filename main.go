package main

import (
	"todoapp-api/controller"
	"todoapp-api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDB()
	r := gin.Default()
	r.GET("/tasks", controller.GetTasks)
	r.Run()
}
