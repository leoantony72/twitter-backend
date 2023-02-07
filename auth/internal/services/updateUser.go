package services

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *userUseCase) Update(user model.User, prevUsername string) error {
	err := u.userRepo.Update(user, prevUsername)
	if err != nil {
		return err
	}
	return nil
}
