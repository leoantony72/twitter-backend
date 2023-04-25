package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

func (u *UserPostgresRepo) Logout(id string) error {
	user := model.User{}
	result := u.db.Model(&user).Where("id =?", id).Update("token","")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
