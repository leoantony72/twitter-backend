package repositories

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"gorm.io/gorm"
)

func (t *TweetRepo) ReTweet(id, user string) error {
	retweet := model.Retweet{}
	retweet.TweetId = id
	retweet.Username = user
	tweet := model.Tweets{}
	result := t.db.Model(&retweet).Create(&retweet)
	if result.Error != nil {
		return result.Error
	}
	result = t.db.Model(&tweet).Where("id=?", id).Update("retweet_count", gorm.Expr("retweet_count + 1"))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
