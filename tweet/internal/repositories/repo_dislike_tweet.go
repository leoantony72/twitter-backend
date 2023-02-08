package repositories

import (
	"errors"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"gorm.io/gorm"
)

func (t *TweetRepo) DislikeTweet(like model.Like) error {
	redis_tweet_like_key := "tweets:" + like.TweetId + ":like"
	redis_tweet_key := "tweets:" + like.TweetId
	
	result := t.db.Model(&like).Where("tweet_id=? AND username=?", like.TweetId, like.Username).Delete(like.TweetId)
	if result.RowsAffected == 0 {
		return errors.New("you have not liked the tweet")
	}
	if result.Error != nil {
		return result.Error
	}
	tweet := model.Tweets{}
	result = t.db.Model(&tweet).Where("id=?", like.TweetId).Update("like_count", gorm.Expr("like_count - 1"))
	if result.Error != nil {
		return result.Error
	}
	t.redis.ZRem(ctx, redis_tweet_like_key, like.Username)
	t.redis.HIncrBy(ctx, redis_tweet_key, "like_count", -1)
	return nil
}
