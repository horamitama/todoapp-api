package handler

import "github.com/gin-gonic/gin"

type TaskHandler struct{}

type TaskHandlerInterfase interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

func NewTaskHandler() TaskHandlerInterfase {
	return &TaskHandler{}
}

func (th *TaskHandler) Create(c *gin.Context) {}

func (th *TaskHandler) Update(c *gin.Context) {}

func (th *TaskHandler) List(c *gin.Context) {}

func (th *TaskHandler) Get(c *gin.Context) {}

func (th *TaskHandler) Delete(c *gin.Context) {}
