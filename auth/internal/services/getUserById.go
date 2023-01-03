package services

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *userUseCase) GetUserbyId(id string) (*model.User, error) {
	user, err := u.userRepo.GetUserbyId(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
