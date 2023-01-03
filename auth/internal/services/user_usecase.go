package services

import (
	"github.com/leoantony72/twitter-backend/auth/internal/ports"
)

type userUseCase struct {
	userRepo ports.UserRepository
}

func NewUseCase(repo ports.UserRepository) ports.UserUseCase {
	return &userUseCase{userRepo: repo}
}



