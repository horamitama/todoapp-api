package handler

import (
	"net/http"
	"time"
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

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
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
	c.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "/", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func (uh *UserHandler) LogOut(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", "/", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
