package services

import (
	// "time"

	"errors"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"github.com/leoantony72/twitter-backend/auth/internal/ports"
)

type userUseCase struct {
	userRepo ports.UserRepository
}

func NewUseCase(repo ports.UserRepository) ports.UserUseCase {
	return &userUseCase{userRepo: repo}
}

func (u *userUseCase) Create(user model.User) error {
	err := u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
func (u *userUseCase) Delete(id string) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (u *userUseCase) GetUserbyId(id string) (*model.User, error) {
	user, err := u.userRepo.GetUserbyId(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userUseCase) Update(user model.User) error {
	err := u.userRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}
func (u *userUseCase) Login(username string, password string) error {
	user, err := u.userRepo.Login(username)
	if err != nil {
		return err
	}
	//password must be hashed before comparing
	user_password := user.Password
	if user_password != password {
		return errors.New("incorrect password or username")
	}
	return nil
}
