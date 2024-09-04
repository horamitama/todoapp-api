package repository

import (
	"todoapp-api/db"
	"todoapp-api/entity"
)

func CreateTask(task *entity.Task) error {
	db := db.NewDB()
	result := db.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTasks(task *entity.Task) error {
	db := db.NewDB()
}
