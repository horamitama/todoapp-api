package main

import (
	"todoapp-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	uh := handler.NewUserHandler()
	r.POST("/signup", uh.SignUpHandler)
	r.POST("/login", uh.LogInHandler)
	r.POST("/logout", uh.LogOutHandler)

	th := handler.NewTaskHandler()
	r.POST("/task", th.CreateTaskHandler)
	r.GET("/task/:userId", th.ListTaskHandler)
	r.GET("/task/:taskId", th.GetTaskHandler)
	r.PUT("/task/:taskId", th.UpdateTaskHandler)
	r.DELETE("/task/:taskId", th.DeleteTaskHandler)
	r.Run()
}
