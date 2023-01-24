package repositories

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"gorm.io/gorm"
)

func (t *TweetRepo) DeleteReTweet(id, user string) error {
	retweet := model.Retweet{}
	tweet := model.Tweets{}
	result := t.db.Model(&retweet).Where("tweet_id =? AND username=?", id, user).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	result = t.db.Model(&tweet).Where("id=?", id).Update("retweet_count", gorm.Expr("retweet_count - 1"))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
