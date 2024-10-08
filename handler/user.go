package handler

import (
	"net/http"
	"todoapp-api/model"
	"todoapp-api/usecase"

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
		c.JSON(http.StatusInternalServerError, gin.H{"err": "request invalid"})
		return
	}

	err = uh.uu.SignUp(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "created user"})
}

func (uh *UserHandler) LogIn(c *gin.Context) {
	var user model.User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "request invalid"})
		return
	}

	token, err := uh.uu.LogIn(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uh *UserHandler) LogOut(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", "/", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
