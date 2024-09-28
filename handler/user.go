package handler

import (
	"net/http"
	"todoapp-api/model"
	usecase "todoapp-api/usercase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uu usecase.UserUsecaseInterface
}

type UserHandlerInterface interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	LogOut(c *gin.Context)
}

func NewUserHandler(uu usecase.UserUsecaseInterface) UserHandlerInterface {
	return &UserHandler{uu}
}

func (uh *UserHandler) SignUp(c *gin.Context) {
	var user model.User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}
}

func (uh *UserHandler) LogIn(c *gin.Context) {}

func (uh *UserHandler) LogOut(c *gin.Context) {}
