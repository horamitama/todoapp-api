package usecase

import (
	"todoapp-api/model"
	"todoapp-api/repository"
	"todoapp-api/util"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	ur repository.UserRepositoryInterface
}

type UserUsecaseInterface interface {
	SignUp(user *model.User) error
	LogIn(user *model.User) (string, error)
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
	err = uu.ur.CreateUser(&newUser)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) LogIn(user *model.User) (string, error) {
	storedUser, err := uu.ur.GetUserByEmail(user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token, err := util.GenerateJwtToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
