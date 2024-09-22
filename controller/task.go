package controller

import (
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
		c.JSON(http.StatusBadRequest, err)
	}
	storedUser := entity.User{}
	err = db.Where("id=?", task.UserRefer).First(&storedUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	newTask := entity.Task{Title: task.Title, Detail: task.Detail, Status: task.Status}
	newTask.UserRefer = storedUser.ID
	err = db.Create(&newTask).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, gin.H{"task": task})
}

func GetTasks(c *gin.Context) {
	// db := db.NewDB()
	user := entity.User{}
	c.Bind(&user)
	c.JSON(200, user)
	// err := c.Bind(&user)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err)
	// }
	// storedUser := entity.User{}
	// err = db.Where("id=?", user.ID).First(&storedUser).Error
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, err)
	// }
	// storedTasks := []entity.Task{}
	// err = db.Where("user_refer=?", storedUser.ID).Find(&storedTasks).Error
	// if err != nil {
	// 	c.JSON(http.StatusNoContent, gin.H{"message": "no content"})
	// }
	// c.JSON(http.StatusOK, gin.H{"tasks": storedTasks})
}
