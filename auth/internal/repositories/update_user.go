package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Update(user model.User) error {
	result := u.db.UpdateColumns(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
