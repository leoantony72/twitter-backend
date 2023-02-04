package repositories

import (
	"fmt"
	"time"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"github.com/redis/go-redis/v9"
)

func (u *UserPostgresRepo) GetUserbyUsername(username string) (*model.User, error) {
	user := model.User{}
	user.Username = username

	//cached data
	redis_key := "users:" + user.Username + ":metadata"
	ok := DoesKeyExist(u, redis_key)
	err := u.redis.HGetAll(ctx, redis_key).Scan(&user)
	if err == nil && err != redis.Nil && ok {
		fmt.Println("cache: ", user)
		fmt.Println("CACHED DATA")
		//decode the time string from redis
		//store it in the User.Date_created struct
		var tm time.Time
		tm.UnmarshalText([]byte(user.Encoded_Date))
		user.Date_created = tm
		return &user, err
	}
	fmt.Println(err)
	fmt.Println("DATABASE DATA")
	result := u.db.Model(&user).Select("id", "username", "email", "date_created", "followers", "following").Where("username =?", username).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	//encode the data
	//store it in User.Encoded_Date Struct
	encoded_date, _ := user.Date_created.MarshalText()
	user.Encoded_Date = string(encoded_date)
	u.redis.HSet(ctx, redis_key, &user)
	u.redis.ExpireAt(ctx, redis_key, time.Now().Add(time.Second*20))
	return &user, nil
}
