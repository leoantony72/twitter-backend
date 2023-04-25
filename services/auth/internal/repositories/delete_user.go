package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Delete(username string) error {
	user := model.User{}
	result := u.db.Model(&user).Where("username =?", username).Delete(username)
	if result.Error != nil {
		return result.Error
	}
	redis_key := "users:" + username + ":*"
	iter := u.redis.Scan(ctx, 0, redis_key, 0).Iterator()
	for iter.Next(ctx) {
		// fmt.Println("keys", iter.Val())
		u.redis.Del(ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return nil
}
