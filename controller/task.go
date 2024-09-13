package controller

import (
	"log"
	"net/http"
	"todoapp-api/db"
	"todoapp-api/entity"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	db := db.NewDB()
	task := entity.Task{}
	err := c.Bind(&task)
	if err != nil {
		log.Fatal("error:", err)
	}
	storedUser := entity.User{}
	err = db.Where("id = ?", task.UserRefer).First(&storedUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	newTask := entity.Task{Title: task.Title, Detail: task.Detail, Status: task.Status}
	newTask.UserRefer = storedUser.ID
	err = db.Create(&newTask).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, gin.H{"task": task})
}

func GetTasks(c *gin.Context) {

}
