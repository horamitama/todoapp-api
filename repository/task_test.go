package repository

import (
	"testing"
	"todoapp-api/model"
	"todoapp-api/util"
)

func TestCreateTask(t *testing.T) {
	testDB := util.SetupTestDB()
	defer util.TeardownTestDB(testDB)
	tr := NewTaskRepository(testDB)

	t.Run("create task", func(t *testing.T) {
		task := model.Task{
			Title:  "hoge",
			Detail: "hoge",
			UserID: 1,
		}
		createdTask, err := tr.CreateTask(&task)
		if err != nil {
			t.Fatal(err)
		}
		if createdTask.Title != task.Title || createdTask.Detail != task.Detail {
			t.Fatal("create user failed")
		}
	})
}

func TestUpdateTask(t *testing.T) {}

func TestFindByUserIDAndTaskID(t *testing.T) {}

func TestFindAllByUserIDTask(t *testing.T) {}

func TestDeleteTask(t *testing.T) {}

// CreateTask(task *model.Task) error
// UpdateTask(task *model.Task, taskID uint) error
// FindByUserIDAndTaskID(task *model.Task, userID uint, taskId uint) error
// FindAllByUserId(task *[]model.Task, userId uint) error
// DeleteTask(task *model.Task, taskID uint) error
