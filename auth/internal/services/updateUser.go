package services

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *userUseCase) Update(user model.User) error {
	err := u.userRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}
