package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Delete(id string) error {
	user := model.User{}
	result := u.db.Model(&user).Where("id =?", id).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
