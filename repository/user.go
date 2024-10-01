package repository

import (
	"todoapp-api/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	CreateUser(user *model.User) error
	GetUserByEmail(user *model.User) (model.User, error)
	LogOut(user *model.User) error
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	err := ur.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByEmail(user *model.User) (model.User, error) {
	var storedUser model.User
	err := ur.db.Where("email = ?", user.Email).Find(&storedUser).Error
	if err != nil {
		return storedUser, err
	}
	return storedUser, nil
}

func (ur *UserRepository) LogOut(user *model.User) error {
	return nil
}
