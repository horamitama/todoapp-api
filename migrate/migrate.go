package main

import (
	"fmt"
	"log"
	"todoapp-api/db"
	"todoapp-api/entity"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&entity.User{}, &entity.Task{})
	if err != nil {
		log.Fatal("Migrate failed")
	}
}
