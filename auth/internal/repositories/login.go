package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Login(username string) (*model.User, error) {
	user := model.User{}
	result := u.db.Select("id","username","password")
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
