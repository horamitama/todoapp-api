package controller

import "github.com/gin-gonic/gin"

func GetTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get tasks",
	})
}

// r.GET("/", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// })
