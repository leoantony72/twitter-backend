package repositories

import (
	"errors"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (t *TweetRepo) LikeTweet(like model.Like) error {
	redis_tweet_key := "tweets:" + like.TweetId
	redis_tweet_like_key := "tweets:" + like.TweetId + ":like"

	result := t.db.Model(&like).Create(&like)
	if result.RowsAffected == 0 {
		return errors.New("invalid Tweet ID")
	}
	if result.Error != nil {
		return result.Error
	}
	tweet := model.Tweets{}
	result = t.db.Model(&tweet).Where("id=? ", like.TweetId).Update("like_count", gorm.Expr("like_count + 1"))
	if result.Error != nil {
		return result.Error
	}
	t.redis.ZAdd(ctx, redis_tweet_like_key, redis.Z{Score: 0, Member: like.Username})
	t.redis.HIncrBy(ctx, redis_tweet_key, "like_count", 1)
	return nil
}
