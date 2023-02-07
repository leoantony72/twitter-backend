package repositories

import (
	"fmt"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

func (u *UserPostgresRepo) Update(user model.User, prevUsername string) error {
	prevUserData := model.User{}
	// var username string
	u.db.Model(&user).Select("username").Where("id=?", user.Id).Scan(&prevUserData)
	// defer user_data.Close()
	result := u.db.Model(&user).Where("id =?", user.Id).Update("username", user.Username)
	if result.Error != nil {
		return result.Error
	}
	redis_key := "users:" + prevUserData.Username + ":*"
	iter := u.redis.Scan(ctx, 0, redis_key, 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
		u.redis.Del(ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return nil
}
