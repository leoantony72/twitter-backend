package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) GetUserbyId(id string) (*model.User, error) {
	user := model.User{}
	// result := u.db.Find(&user, id)
	result := u.db.Select("id","username","date_created")
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
