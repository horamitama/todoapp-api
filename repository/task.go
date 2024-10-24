package repository

import (
	"todoapp-api/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

type TaskRepositoryInterfase interface {
	CreateTask(task *model.Task) (*model.Task, error)
	UpdateTask(task *model.Task, taskID uint) error
	FindByUserIDAndTaskID(task *model.Task, userID uint, taskId uint) error
	FindAllTasksByUserID(task *[]model.Task, userId uint) error
	DeleteTask(task *model.Task, taskID uint) error
}

func NewTaskRepository(db *gorm.DB) TaskRepositoryInterfase {
	return &TaskRepository{db}
}

func (tr *TaskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	err := tr.db.Create(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *TaskRepository) FindByUserIDAndTaskID(task *model.Task, userID uint, taskID uint) error {
	err := tr.db.Where("user_id = ? AND id = ?", userID, taskID).Find(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) FindAllTasksByUserID(tasks *[]model.Task, userID uint) error {
	err := tr.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) UpdateTask(task *model.Task, taskID uint) error {
	err := tr.db.Where("id = ?", taskID).Save(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteTask(task *model.Task, taskID uint) error {
	err := tr.db.Where("id = ?", taskID).Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}
