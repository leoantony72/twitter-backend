package services

import (
	"time"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"github.com/leoantony72/twitter-backend/auth/internal/utils"
)

func (u *userUseCase) Create(user model.User) error {

	HashedPassword, salt, err := utils.GenerateHashfromPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = HashedPassword
	user.Salt = salt
	user.Id = utils.GenerateID().String()
	user.Date_created = time.Now()
	err = u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
