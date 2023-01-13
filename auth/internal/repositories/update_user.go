package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Update(user model.User) error {
	result := u.db.Model(&user).Where("id =?", user.Id).Update("username", user.Username)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
