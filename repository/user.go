package repository

import (
	"todoapp-api/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	SignUp(user *model.User) error
	LogIn(user *model.User) error
	LogOut(user *model.User) error
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (ur *UserRepository) SignUp(user *model.User) error {
	err := ur.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) LogIn(user *model.User) error {
	return nil
}

func (ur *UserRepository) LogOut(user *model.User) error {
	return nil
}
