package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

type UserHandlerInterface interface {
	SignUpHandler(c *gin.Context)
	LogInHandler(c *gin.Context)
	LogOutHandler(c *gin.Context)
}

func NewUserHandler() UserHandlerInterface {
	return &UserHandler{}
}

func (uh *UserHandler) SignUpHandler(c *gin.Context) {}

func (uh *UserHandler) LogInHandler(c *gin.Context) {}

func (uh *UserHandler) LogOutHandler(c *gin.Context) {}
