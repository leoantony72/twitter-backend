package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (u *UserPostgresRepo) FollowUser(follow model.User_followers) error {
	user := model.User{}
	err := u.db.Model(&follow).Create(&follow)
	if err.Error != nil {
		return err.Error
	}
	//update follower -> following count +1
	u.db.Model(&user).Where("username", follow.Follower).Update("following_count", gorm.Expr("following_count + 1"))

	//update followee - followers count +1
	u.db.Model(&user).Where("username", follow.Followee).Update("follower_count", gorm.Expr("follower_count + 1"))

	//@Add followee to client's followings cache
	//@Add client to followee's followers cache
	lredis_client_key := "users:" + follow.Follower + ":following"
	lredis_followee_key := "users:" + follow.Followee + ":followers"
	u.redis.ZAdd(ctx, lredis_client_key, redis.Z{Member: follow.Followee})
	u.redis.ZAdd(ctx, lredis_followee_key, redis.Z{Member: follow.Follower})

	//@Update the client's following count
	//@Update the followee's followeres count
	redis_client_key := "users:" + follow.Follower + ":following_count"
	redis_followee_key := "users:" + follow.Followee + ":follower_count"
	u.redis.IncrBy(ctx, redis_client_key, 1)
	u.redis.IncrBy(ctx, redis_followee_key, 1)
	// u.redis.ExpireAt(ctx, hredis_client_key, time.Now().Add(time.Second*10))
	// u.redis.HIncrBy(ctx, hredis_followee_key, "followers", 1)
	// u.redis.ExpireAt(ctx, hredis_followee_key, time.Now().Add(time.Second*10))

	return err.Error
}
