package controller

import (
	"net/http"
	"todoapp-api/db"
	"todoapp-api/entity"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpController(c *gin.Context) {
	db := db.NewDB()
	user := entity.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	newUser := entity.User{Name: user.Name, Email: user.Email, Password: string(hash)}
	err = db.Create(&newUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(200, user)
}
