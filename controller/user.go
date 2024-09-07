package controller

import (
	"net/http"
	"os"
	"time"
	"todoapp-api/db"
	"todoapp-api/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func SigninController(c *gin.Context) {
	db := db.NewDB()
	user := entity.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	storedUser := entity.User{}
	err = db.Where("email = ?", user.Email).Find(&storedUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.SetCookie("token", tokenString, 3600, "", os.Getenv("API_DOMAIN"), true, true)
}
