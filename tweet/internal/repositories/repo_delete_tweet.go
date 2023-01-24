package repositories

import (
	"errors"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetRepo) DeleteTweet(id, user string) error {
	tweet := model.Tweets{}
	result := t.db.Model(&tweet).Where("id=? AND username=?", id, user).Delete(id)
	if result.RowsAffected == 0 {
		return errors.New("invalid Command")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
