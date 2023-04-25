package repositories

import (
	"errors"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (t *TweetRepo) ReTweet(retweet model.Retweet) error {
	redis_tweet_key := "tweets:" + retweet.TweetId
	redis_tweet_retweet_key := "tweets:" + retweet.TweetId + ":retweet"
	tweet := model.Tweets{}
	result := t.db.Model(&retweet).Create(&retweet)
	if result.RowsAffected == 0 {
		return errors.New("invalid Tweet ID")
	}
	if result.Error != nil {
		return result.Error
	}
	result = t.db.Model(&tweet).Where("id=?", retweet.TweetId).Update("retweet_count", gorm.Expr("retweet_count + 1"))
	if result.Error != nil {
		return result.Error
	}
	t.redis.HIncrBy(ctx, redis_tweet_key, "retweet_count", 1)
	t.redis.ZAdd(ctx, redis_tweet_retweet_key, redis.Z{Score: 0, Member: retweet.Username})
	return nil
}
