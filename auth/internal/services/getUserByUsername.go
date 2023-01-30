package services

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *userUseCase) GetUserbyUsername(username string) (*model.User, error) {
	user, err := u.userRepo.GetUserbyUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
