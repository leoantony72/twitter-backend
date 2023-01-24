package repositories

import (
	"errors"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"gorm.io/gorm"
)

func (t *TweetRepo) DeleteReTweet(retweet model.Retweet) error {
	// retweet := model.Retweet{}
	tweet := model.Tweets{}
	result := t.db.Model(&retweet).Where("tweet_id =? AND username=?", retweet.TweetId, retweet.Username).Delete(retweet.TweetId)
	if result.RowsAffected == 0 {
		return errors.New("you have not retweeted")
	}
	if result.Error != nil {
		return result.Error
	}
	result = t.db.Model(&tweet).Where("id=?", retweet.TweetId).Update("retweet_count", gorm.Expr("retweet_count - 1"))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
