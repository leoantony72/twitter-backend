package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetRepo) GetTweetById(id string) (*model.Tweets, error) {
	tweet := model.Tweets{}
	redis_key := "tweets:" + id
	ok := DoesKeyExist(t, redis_key)
	err := t.redis.HGetAll(ctx, redis_key).Scan(&tweet)
	fmt.Println("cached data", ok, err)
	if ok && err == nil {
		var tm time.Time
		tm.UnmarshalText([]byte(tweet.Encoded_date))
		tweet.CreatedAt = tm
		return &tweet, nil
	}
	result := t.db.Model(&tweet).Where("id=?", id).Scan(&tweet)
	if result.RowsAffected == 0 {
		return nil, errors.New("invalid Tweet ID")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	encoded_date, _ := tweet.CreatedAt.MarshalText()
	tweet.Encoded_date = string(encoded_date)
	t.redis.HSet(ctx, redis_key, &tweet)
	t.redis.ExpireAt(ctx, redis_key, time.Now().Add(time.Second*20))
	return &tweet, nil
}
