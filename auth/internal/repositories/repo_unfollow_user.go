package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"gorm.io/gorm"
)

func (u *UserPostgresRepo) UnfollowUser(follow model.User_followers) error {
	user := model.User{}
	result := u.db.Model(&follow).Where("follower = ? AND followee = ?", follow.Follower, follow.Followee).Delete(&follow)
	if result.Error != nil {
		return result.Error
	}

	//update follower -> following count -1
	u.db.Model(&user).Where("username", follow.Follower).Update("following_count", gorm.Expr("following_count - 1"))

	//update followee - followers count -1
	u.db.Model(&user).Where("username", follow.Followee).Update("follower_count", gorm.Expr("follower_count - 1"))

	//@Remove followee to client's followings cache
	//@Remove client to followee's followers cache
	lredis_client_key := "users:" + follow.Follower + ":following"
	lredis_followee_key := "users:" + follow.Followee + ":followers"
	// u.redis.ZAdd(ctx, lredis_client_key, redis.Z{Member: follow.Followee})
	u.redis.ZRem(ctx, lredis_client_key, follow.Followee)
	u.redis.ZRem(ctx, lredis_followee_key, follow.Follower)

	//@Update the client's following count-decrease
	//@Update the followee's followeres count-decrease
	redis_client_key := "users:" + follow.Follower + ":metadata"
	redis_followee_key := "users:" + follow.Followee + ":metadata"
	u.redis.DecrBy(ctx, redis_client_key, 1)
	u.redis.DecrBy(ctx, redis_followee_key, 1)
	return nil
}
