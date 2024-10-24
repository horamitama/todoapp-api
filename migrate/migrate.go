package main

import (
	"fmt"
	"log"
	"todoapp-api/model"
	"todoapp-api/util"
)

func main() {
	dbConn := util.NewDB()
	defer fmt.Println("Successfully migrated")
	defer util.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatal("Migrate failed")
	}
}
