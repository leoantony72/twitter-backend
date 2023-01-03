package services

import "errors"

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
