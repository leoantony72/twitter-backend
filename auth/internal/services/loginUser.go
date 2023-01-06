package services

import (
	"errors"
	"github.com/leoantony72/twitter-backend/auth/internal/utils"
)

func (u *userUseCase) Login(username string, password string) error {
	//get data from db store in userData
	userData, err := u.userRepo.Login(username)
	if err != nil {
		return err
	}
	//call the function VerifyPassword from utils
	//It returns a boolean value and err
	//is it's true user credential is valid
	match, _ := utils.VerifyPassword(password, userData.Password, userData.Salt)
	if !match {
		return errors.New("incorrect username or password")
	}

	//Generate JWT and send it to user

	return nil
}
