package usecase

import (
	"todoapp-api/model"
)

type UserUsecase struct{}

type UserUsecaseInterface interface {
	SignUp(user *model.User) error
	LogIn(user *model.User) error
	LogOut(user *model.User) error
}

func NewUserUsecase() UserUsecaseInterface {
	return &UserUsecase{}
}

func (uu *UserUsecase) SignUp(user *model.User) error {
	return nil
}

func (uu *UserUsecase) LogIn(user *model.User) error {
	return nil
}

func (uu *UserUsecase) LogOut(user *model.User) error {
	return nil
}
