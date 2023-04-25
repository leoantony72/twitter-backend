package repositories

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

type TempUserFollows struct {
	Followers []string `redis:"followers"`
	Following []string `redis:"following"`
}

func (u *UserPostgresRepo) GetUserbyUsername(username string) (*model.User, error) {
	user := model.User{}
	userFollow := []model.User_followers{}
	TempUserFollows := TempUserFollows{}
	user.Username = username

	//@Redis Key
	redis_key := "users:" + user.Username + ":metadata"
	redisFollowingKey := "users:" + user.Username + ":following"
	redisFollowersKey := "users:" + user.Username + ":followers"
	redisFollowerCount := "users:" + user.Username + ":follower_count"
	redisFollowingCount := "users:" + user.Username + ":following_count"

	ok := DoesKeyExist(u, redis_key)
	err := u.redis.HGetAll(ctx, redis_key).Scan(&user)

	//@Get Followers from Redis Sorted Set
	//@Get Followings from Redis Sorted Set
	u.redis.ZRange(ctx, redisFollowersKey, 0, -1).ScanSlice(&TempUserFollows.Followers)
	u.redis.ZRange(ctx, redisFollowingKey, 0, -1).ScanSlice(&TempUserFollows.Following)
	Followercount, _ := u.redis.Get(ctx, redisFollowerCount).Result()
	Followingcount, _ := u.redis.Get(ctx, redisFollowingCount).Result()
	if err == nil && err != redis.Nil && ok {
		user.Followers = TempUserFollows.Followers
		user.Following = TempUserFollows.Following
		user.Follower_Count, _ = strconv.Atoi(Followercount)
		user.Following_Count, _ = strconv.Atoi(Followingcount)
		// fmt.Println("CACHED DATA")
		//@decode the time string from redis
		//@store it in the User.Date_created struct
		var tm time.Time
		tm.UnmarshalText([]byte(user.Encoded_Date))
		user.Date_created = tm
		return &user, err
	}
	// fmt.Println(err)
	// fmt.Println("DATABASE DATA")
	user_result := u.db.Model(&user).Select("id", "username", "email", "date_created", "follower_count", "following_count").Where("username =?", username).Scan(&user)
	//@Check if user exist
	if user.Id == "" {
		return nil, errors.New("user does not exist")
	}
	if user_result.Error != nil {
		return nil, user_result.Error
	}
	//@Get User followers
	//@Get User followings
	following_rows, _ := u.db.Model(&userFollow).Select("followee").Where("follower=?", user.Username).Rows()
	followers_rows, _ := u.db.Model(&userFollow).Select("follower").Where("followee=?", user.Username).Rows()
	defer following_rows.Close()
	defer followers_rows.Close()
	for following_rows.Next() {
		u.db.ScanRows(following_rows, &TempUserFollows.Following)
	}
	for followers_rows.Next() {
		u.db.ScanRows(followers_rows, &TempUserFollows.Followers)
	}

	//@encode the data
	//@store it in User.Encoded_Date Struct
	encoded_date, _ := user.Date_created.MarshalText()
	user.Encoded_Date = string(encoded_date)
	rerr := u.redis.HSet(ctx, redis_key, &user)
	u.redis.ExpireAt(ctx, redis_key, time.Now().Add(time.Second*20))
	fmt.Println("redis Err: ", rerr.Err())
	user.Followers = TempUserFollows.Followers
	user.Following = TempUserFollows.Following

	u.redis.Set(ctx, redisFollowerCount, user.Follower_Count, 0)
	u.redis.Set(ctx, redisFollowingCount, user.Following_Count, 0)
	for _, user := range TempUserFollows.Followers {
		u.redis.ZAddNX(ctx, redisFollowersKey, redis.Z{Member: user})
	}

	for _, user := range TempUserFollows.Following {
		u.redis.ZAddNX(ctx, redisFollowingKey, redis.Z{Member: user})
	}
	return &user, nil
}
