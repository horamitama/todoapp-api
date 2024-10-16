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
	Update(task *model.Task) error
	FindByUserIDAndID(task *model.Task, userID uint, taskId uint) error
	FindAllByUserId(task *[]model.Task, userId uint) error
	Delete(task *model.Task) error
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

func (tr *TaskRepository) FindByUserIDAndID(task *model.Task, userID uint, taskID uint) error {
	err := tr.db.Where("user_id = ? AND id = ?", userID, taskID).Find(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) FindAllByUserId(tasks *[]model.Task, userID uint) error {
	err := tr.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) Update(task *model.Task) error {
	err := tr.db.Save(task).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) Delete(task *model.Task) error {
	err := tr.db.Delete(task).Error
	if err != nil {
		return err
	}
	return nil
}
