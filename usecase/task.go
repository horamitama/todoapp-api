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
	Update(task *model.Task, taskID uint) error
	Delete(task *model.Task, taskID uint) error
}

func NewTaskUsecase(tr repository.TaskRepositoryInterfase) TaskUsecaseInterface {
	return &TaskUsecase{tr}
}

func (tu *TaskUsecase) Create(task *model.Task) error {
	err := tu.tr.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskUsecase) List(task *model.Task) (*[]model.Task, error) {
	var storedTasks []model.Task
	err := tu.tr.FindAllTasksByUserID(&storedTasks, task.UserID)
	if err != nil {
		return nil, err
	}
	return &storedTasks, nil
}

func (tu *TaskUsecase) Get(task *model.Task, userID uint, taskID uint) (*model.Task, error) {
	var storedTask model.Task

	err := tu.tr.FindByUserIDAndTaskID(&storedTask, uint(userID), uint(taskID))
	if err != nil {
		return nil, err
	}

	return &storedTask, nil
}

func (tu *TaskUsecase) Update(task *model.Task, taskID uint) error {
	err := tu.tr.UpdateTask(task, taskID)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskUsecase) Delete(task *model.Task, taskID uint) error {
	err := tu.tr.DeleteTask(task, taskID)
	if err != nil {
		return err
	}
	return nil
}
