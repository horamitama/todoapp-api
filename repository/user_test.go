package repository

import (
	"testing"
	"todoapp-api/model"
	"todoapp-api/util"
)

func TestCreateUser(t *testing.T) {
	testDB := util.SetupTestDB()
	defer util.TeardownTestDB(testDB)
	ur := NewUserRepository(testDB)

	t.Run("create user", func(t *testing.T) {
		user := model.User{
			Email:    "hoge@a.com",
			Password: "hogehoge",
		}
		createdUser, err := ur.CreateUser(&user)
		if err != nil {
			t.Fatal(err)
		}
		if createdUser.Email != user.Email {
			t.Fatal("create user failed")
		}
	})
}

func TestGetUserByEmail(t *testing.T) {
	testDB := util.SetupTestDB()
	defer util.TeardownTestDB(testDB)
	ur := NewUserRepository(testDB)

	t.Run("get user by email", func(t *testing.T) {
		email := "hoge@a.com"
		password := "hogehoge"
		user := model.User{
			Email:    email,
			Password: password,
		}
		_, err := ur.CreateUser(&user)
		if err != nil {
			t.Fatal(err)
		}
		var storedUser model.User
		err = ur.GetUserByEmail(&storedUser, email)
		if err != nil {
			t.Fatal(err)
		}
	})
}
