package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserControllerInterface interface {
	SignUp(c *gin.Context) error
	LogIn(c *gin.Context) (string, error)
	LogOut(c *gin.Context) error
}

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) UserControllerInterface {
	return &UserController{db}
}

func (uc *UserController) SignUp(c *gin.Context) error {
	return nil
}

func (uc *UserController) LogIn(c *gin.Context) (string, error) {
	return "", nil
}

func (uc *UserController) LogOut(c *gin.Context) error {
	return nil
}

// func SignUp(c *gin.Context) {
// 	db := db.NewDB()
// 	user := model.User{}
// 	err := c.Bind(&user)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	newUser := model.User{Name: user.Name, Email: user.Email, Password: string(hash)}
// 	err = db.Create(&newUser).Error
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	c.JSON(200, user)
// }
//
// func SignIn(c *gin.Context) {
// 	db := db.NewDB()
// 	user := model.User{}
// 	err := c.Bind(&user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 	}
// 	storedUser := model.User{}
// 	err = db.Where("email = ?", user.Email).Find(&storedUser).Error
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 	}
//
// 	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 	}
//
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": user.ID,
// 		"exp":     time.Now().Add(time.Hour * 12).Unix(),
// 	})
// 	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 	}
//
// 	c.SetCookie("token", tokenString, 3600, "", os.Getenv("API_DOMAIN"), true, true)
// }
//
// func SignOut(c *gin.Context) {
// 	user := model.User{}
// 	err := c.Bind(&user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 	}
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 	}
// 	c.SetCookie("token", "", 0, "", os.Getenv("API_DOMAIN"), false, false)
// }
