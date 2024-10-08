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

	// user route
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	r.POST("/signup", userHandler.SignUp)
	r.POST("/login", userHandler.LogIn)
	r.POST("/logout", userHandler.LogOut)

	// task route
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUseCase)
	r.POST("/task", taskHandler.Create)
	r.GET("/task", taskHandler.List)
	r.GET("/task/:ID", taskHandler.Get)
	r.PUT("/task/:ID", taskHandler.Update)
	r.DELETE("/task/:ID", taskHandler.Delete)

	r.Run(":8080")
}
