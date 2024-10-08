package repository

import (
	"todoapp-api/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

type TaskRepositoryInterfase interface {
	Create(task *model.Task) error
	FindById(taskId int) error
	FindAllByUserId(task *[]model.Task, userId uint) error
}

func NewTaskRepository(db *gorm.DB) TaskRepositoryInterfase {
	return &TaskRepository{db}
}

func (tr *TaskRepository) Create(task *model.Task) error {
	err := tr.db.Create(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) FindById(taskId int) error {
	return nil
}

func (tr *TaskRepository) FindAllByUserId(tasks *[]model.Task, userID uint) error {
	err := tr.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return err
	}
	return nil
}
