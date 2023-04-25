package services

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *userUseCase) GetToken(refreshToken string) (model.User, error) {
	user, err := u.userRepo.GetToken(refreshToken)

	if err != nil {
		return user, err
	}
	return user, nil
}
