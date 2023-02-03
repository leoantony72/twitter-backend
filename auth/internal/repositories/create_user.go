package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

func (u *UserPostgresRepo) Create(user model.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	redis_key := "users:" + user.Username + ":" + "metadata"
	//Encoded data
	encoded_date, _ := user.Date_created.MarshalText()
	user.Encoded_Date = string(encoded_date)
	_, err := Cache(u, redis_key, "hset", user, "")
	if err != nil {
		return err
	}
	return nil
}
