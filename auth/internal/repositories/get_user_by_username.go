package repositories

import (
	"fmt"
	"time"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

func (u *UserPostgresRepo) GetUserbyUsername(username string) (*model.User, error) {
	user := model.User{}
	user.Username = username

	//cached data
	redis_key := "users:" + user.Username + ":metadata"
	user_cache, err := Cache(u, redis_key, "hgetall", user, "")
	if err == nil {
		//decode the time string from redis
		//store it in the User.Date_created struct
		var tm time.Time
		tm.UnmarshalText([]byte(user_cache.Encoded_Date))
		user_cache.Date_created = tm
		return user_cache, err
	}
	fmt.Println(err)
	result := u.db.Model(&user).Select("id", "username", "email", "date_created", "followers", "following").Where("username =?", username).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	//encode the data
	//store it in User.Encoded_Date Struct
	encoded_date, _ := user.Date_created.MarshalText()
	user.Encoded_Date = string(encoded_date)
	Cache(u, redis_key, "hset", user, "")
	return &user, nil
}
