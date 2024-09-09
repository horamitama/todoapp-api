package controller

import (
	"log"
	"todoapp-api/db"
	"todoapp-api/entity"

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
	c.JSON(200, gin.H{"task": task})
}

// func UpdateTask(c *gin.Context) {
// 	db := db.NewDB()
// 	task := entity.Task{}
// 	err := c.Bind(&task)
// 	if err != nil {
// 		log.Fatal()
// 	}
// }

func DeleteTask(c *gin.Context) {}
