package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Delete(username string) error {
	user := model.User{}
	result := u.db.Model(&user).Where("username =?", username).Delete(username)
	if result.Error != nil {
		return result.Error
	}
	redis_key := "users:" + username + ":metadata"
	redis_following_key := "users:" + username + ":following"
	redis_followers_key := "users:" + username + ":followers"
	redis_following_count_key := "users:" + username + ":following_count"
	redis_followers_count_key := "users:" + username + ":follower_count"
	u.redis.HDel(ctx, redis_key)
	u.redis.HDel(ctx, redis_followers_key)
	u.redis.HDel(ctx, redis_following_key)
	u.redis.HDel(ctx, redis_following_count_key)
	u.redis.HDel(ctx, redis_followers_count_key)
	return nil
}
