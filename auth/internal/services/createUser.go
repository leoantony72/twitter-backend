package services

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *userUseCase) Create(user model.User) error {
	// password := user.Password
	err := u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
