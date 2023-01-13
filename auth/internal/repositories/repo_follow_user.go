package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) FollowUser(follow model.User_followers) error {
	err := u.db.Model(&follow).Create(&follow)
	return err.Error
}
