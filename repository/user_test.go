package repository

import (
	"fmt"
	"testing"
	"todoapp-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	port := "54320"
	name := "postgres"
	user := "postgres"
	password := "postgres"
	host := "localhost"

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user,
		password, host,
		port, name)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("db connect failed")
	}

	return db
}

func teardownTestDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		panic("db close failed")
	}
}

func TestCreateUser(t *testing.T) {
	testDB := setupTestDB()
	defer teardownTestDB(testDB)
	ur := NewUserRepository(testDB)

	t.Run("create user", func(t *testing.T) {
		user := model.User{
			Email:    "hoge@a.com",
			Password: "hogehoge",
		}
		createdUser, err := ur.CreateUser(&user)
		if err != nil {
			t.Error(err)
		}
		if createdUser.Email != user.Email {
			t.Error("create user failed")
		}
	})
}
