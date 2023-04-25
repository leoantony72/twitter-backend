package services

import "github.com/leoantony72/twitter-backend/auth/internal/utils"

func (u *userUseCase) GenerateAccessToken(name, Id string) (string, error) {
	token, err := utils.GenerateAccessToken(name, Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
