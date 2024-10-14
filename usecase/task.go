package usecase

import (
	"todoapp-api/model"
	"todoapp-api/repository"
)

type TaskUsecase struct {
	tr repository.TaskRepositoryInterfase
}

type TaskUsecaseInterface interface {
	Create(*model.Task) error
	List(*model.Task) (*[]model.Task, error)
	Get(*model.Task) (*model.Task, error)
	Update(*model.Task) error
	Delete(*model.Task) error
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

func (tu *TaskUsecase) Get(task *model.Task) (*model.Task, error) {
	var storedTask model.Task
	err := tu.tr.FindById(storedTask, task.ID)
	if err != nil {
		return nil, err
	}
	return &storedTask, nil
}

func (tu *TaskUsecase) Update(task *model.Task) error {
	err := tu.tr.Update(task)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskUsecase) Delete(task *model.Task) error {
	return nil
}
