package main

import (
	"fmt"
	"todoapp-api/db"
	"todoapp-api/entity"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&entity.User{}, &entity.Task{})
}
