package usecase

import (
	"todoapp-api/model"
	"todoapp-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	ur repository.UserRepositoryInterface
}

type UserUsecaseInterface interface {
	SignUp(user *model.User) error
	LogIn(user *model.User) error
	LogOut(user *model.User) error
}

func NewUserUsecase(ur repository.UserRepositoryInterface) UserUsecaseInterface {
	return &UserUsecase{ur}
}

func (uu *UserUsecase) SignUp(user *model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	newUser := model.User{
		Email:    user.Email,
		Password: string(hash),
	}
	err = uu.ur.SignUp(&newUser)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) LogIn(user *model.User) error {
	return nil
}

func (uu *UserUsecase) LogOut(user *model.User) error {
	return nil
}
