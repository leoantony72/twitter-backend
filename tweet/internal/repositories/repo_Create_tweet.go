package repositories

import (
	// "time"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"github.com/redis/go-redis/v9"
)

func (t *TweetRepo) CreateTweet(tweet model.Tweets) error {

	result := t.db.Create(&tweet)
	if result.Error != nil {
		return result.Error
	}
	redis_key := "tweets:" + tweet.Id
	user_redis_key := "users:" + tweet.Username + ":tweets"
	encoded_date, _ := tweet.CreatedAt.MarshalText()
	tweet.Encoded_date = string(encoded_date)
	t.redis.HSet(ctx, redis_key, &tweet)
	t.redis.ZAdd(ctx, user_redis_key, redis.Z{Score: 0, Member: tweet.Id})
	// t.redis.ExpireAt(ctx, redis_key, time.Now().Add(time.Second*20))
	return nil
}
