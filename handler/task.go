package handler

import (
	"net/http"
	"todoapp-api/model"
	"todoapp-api/usecase"
	"todoapp-api/util"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	tu usecase.TaskUsecaseInterface
}

type TaskHandlerInterfase interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

func NewTaskHandler(tu usecase.TaskUsecaseInterface) TaskHandlerInterfase {
	return &TaskHandler{tu}
}

func (th *TaskHandler) Create(c *gin.Context) {
	var task model.Task
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	authHeader := c.GetHeader("Authorization")
	userID, err := util.GetUserIDFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	NewTask := model.Task{
		Title:  task.Title,
		Detail: task.Detail,
		Status: task.Status,
	}

	NewTask.UserID = userID
	err = th.tu.Create(&NewTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task created"})
}

func (th *TaskHandler) List(c *gin.Context) {
	var task model.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	authHeader := c.GetHeader("Authorization")
	userID, err := util.GetUserIDFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	task.UserID = userID
	storedTask, err := th.tu.List(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, storedTask)
}

func (th *TaskHandler) Get(c *gin.Context) {
	var task model.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	authHeader := c.GetHeader("Authorization")
	userID, err := util.GetUserIDFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	task.UserID = userID
	storedTask, err := th.tu.Get(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, storedTask)
}

func (th *TaskHandler) Update(c *gin.Context) {
	var task model.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	authHeader := c.GetHeader("Authorization")
	userID, err := util.GetUserIDFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	storedTask, err := th.tu.Get(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updateTask := model.Task{
		Title:  task.Title,
		Detail: task.Detail,
		Status: task.Status,
	}

	updateTask.UserID = userID
	err = th.tu.Update(&updateTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": storedTask})
}

func (th *TaskHandler) Delete(c *gin.Context) {
	var task model.Task
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
