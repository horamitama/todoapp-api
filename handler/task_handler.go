package handler

import "github.com/gin-gonic/gin"

type TaskHandler struct{}

type TaskHandlerInterfase interface {
	CreateTaskHandler(c *gin.Context)
	UpdateTaskHandler(c *gin.Context)
	ListTaskHandler(c *gin.Context)
	GetTaskHandler(c *gin.Context)
	DeleteTaskHandler(c *gin.Context)
}

func NewTaskHandler() TaskHandlerInterfase {
	return &TaskHandler{}
}

func (th *TaskHandler) CreateTaskHandler(c *gin.Context) {}

func (th *TaskHandler) UpdateTaskHandler(c *gin.Context) {}

func (th *TaskHandler) ListTaskHandler(c *gin.Context) {}

func (th *TaskHandler) GetTaskHandler(c *gin.Context) {}

func (th *TaskHandler) DeleteTaskHandler(c *gin.Context) {}
