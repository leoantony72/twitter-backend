package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) GetUserbyUsername(username string) (*model.User, error) {
	user := model.User{}
	// result := u.db.Find(&user, id)
	result := u.db.Model(&user).Select("id", "username","email","date_created", "followers", "following").Where("username =?", username).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
