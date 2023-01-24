package repositories

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"gorm.io/gorm"
)

func (t *TweetRepo) DislikeTweet(id, user string) error {
	like := model.Like{}
	like.TweetId = id
	like.Username = user
	result := t.db.Model(&like).Where("id=? AND username=?", like.TweetId, like.Username).Delete(like.TweetId)
	if result.Error != nil {
		return result.Error
	}
	tweet := model.Tweets{}
	result = t.db.Model(&tweet).Where("id=? and username=?", like.TweetId, like.Username).Update("like_count", gorm.Expr("like_count - 1"))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
