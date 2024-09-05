package controller

import (
	"log"
	"todoapp-api/entity"
	"todoapp-api/repository"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get tasks",
	})
}

func CreateTask(c *gin.Context) {
	task := entity.Task{}
	err := c.Bind(&task)
	if err != nil {
		log.Fatal("error:", err)
	}
	err = repository.CreateTask(&task)
	if err != nil {
		log.Fatal()
	}
	c.JSON(200, gin.H{"task": task})
}
