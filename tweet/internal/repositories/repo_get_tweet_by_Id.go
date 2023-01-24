package repositories

import (
	"errors"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetRepo) GetTweetById(id string) (*model.Tweets, error) {
	tweet := model.Tweets{}
	result := t.db.Model(&tweet).Where("id=?", id).Scan(&tweet)
	if result.RowsAffected == 0 {
		return nil, errors.New("invalid Tweet ID")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &tweet, nil
}
