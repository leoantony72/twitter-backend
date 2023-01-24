package repositories

import (
	"errors"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"gorm.io/gorm"
)

func (t *TweetRepo) LikeTweet(like model.Like) error {
	// like := model.Like{}
	// like.TweetId = id
	// like.Username = user
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
	return nil
}
