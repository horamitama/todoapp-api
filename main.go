package main

import (
	"todoapp-api/db"
	"todoapp-api/handler"
	"todoapp-api/repository"
	"todoapp-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := db.NewDB()

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)
	r.POST("/signup", uh.SignUp)
	r.POST("/login", uh.LogIn)
	r.POST("/logout", uh.LogOut)

	// th := handler.NewTaskHandler()
	// r.POST("/task", th.Create)
	// r.GET("/task/:userId", th.List)
	// r.GET("/task/:Id", th.Get)
	// r.PUT("/task/:Id", th.Update)
	// r.DELETE("/task/:Id", th.Delete)
	r.Run()
}
