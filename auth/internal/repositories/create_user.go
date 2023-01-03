package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Create(user model.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
