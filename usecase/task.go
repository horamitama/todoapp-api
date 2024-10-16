package usecase

import (
	"todoapp-api/model"
	"todoapp-api/repository"
)

type TaskUsecase struct {
	tr repository.TaskRepositoryInterfase
}

type TaskUsecaseInterface interface {
	Create(task *model.Task) error
	List(task *model.Task) (*[]model.Task, error)
	Get(task *model.Task, userID uint, taskId uint) (*model.Task, error)
	Update(task *model.Task, taskId uint) error
	Delete(task *model.Task, taskId uint) error
}

func NewTaskUsecase(tr repository.TaskRepositoryInterfase) TaskUsecaseInterface {
	return &TaskUsecase{tr}
}

func (tu *TaskUsecase) Create(task *model.Task) error {
	err := tu.tr.Create(task)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskUsecase) List(task *model.Task) (*[]model.Task, error) {
	var storedTasks []model.Task
	err := tu.tr.FindAllByUserId(&storedTasks, task.UserID)
	if err != nil {
		return nil, err
	}
	return &storedTasks, nil
}

func (tu *TaskUsecase) Get(task *model.Task, userID uint, taskID uint) (*model.Task, error) {
	var storedTask model.Task

	err := tu.tr.FindByUserIDAndID(&storedTask, uint(userID), uint(taskID))
	if err != nil {
		return nil, err
	}

	return &storedTask, nil
}

func (tu *TaskUsecase) Update(task *model.Task, taskId uint) error {
	err := tu.tr.Update(task)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskUsecase) Delete(task *model.Task, taskId uint) error {
	err := tu.tr.Delete(task)
	if err != nil {
		return err
	}
	return nil
}
