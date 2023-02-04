package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Delete(username string) error {
	user := model.User{}
	result := u.db.Model(&user).Where("username =?", username).Delete(username)
	if result.Error != nil {
		return result.Error
	}
	redis_key := "users:" + username + ":metadata"
	u.redis.HDel(ctx, redis_key)
	return nil
}
