package repository

import (
	"todoapp-api/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByEmail(user *model.User, email string) error
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(user *model.User, email string) error {
	err := ur.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}
