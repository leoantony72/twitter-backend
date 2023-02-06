package repositories

import (
	"time"

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
	create_user_cache(u, redis_key, &user)
	return nil
}
func create_user_cache(u *UserPostgresRepo, key string, values *model.User) {
	u.redis.HSet(ctx, key, &values)

	client_follower_key := "users:" + values.Username + ":follower_count"
	client_following_key := "users:" + values.Username + ":following_count"
	u.redis.IncrBy(ctx, client_follower_key, 0)
	u.redis.IncrBy(ctx, client_following_key, 0)
	u.redis.ExpireAt(ctx, key, time.Now().Add(time.Second*10))
}
