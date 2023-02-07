package repositories

import (
	"time"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetRepo) CreateTweet(tweet model.Tweets) error {

	result := t.db.Create(&tweet)
	if result.Error != nil {
		return result.Error
	}
	redis_key := "tweets:" + tweet.Id
	t.redis.HSet(ctx, redis_key, &tweet)
	t.redis.ExpireAt(ctx, redis_key, time.Now().Add(time.Second*60))
	return nil
}
