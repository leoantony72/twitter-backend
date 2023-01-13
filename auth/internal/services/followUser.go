package services

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *userUseCase) FollowUser(follow model.User_followers) error {
	err := u.userRepo.FollowUser(follow)
	return err
}
