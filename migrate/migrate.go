package main

import (
	"fmt"
	"log"
	"todoapp-api/db"
	"todoapp-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatal("Migrate failed")
	}
}
