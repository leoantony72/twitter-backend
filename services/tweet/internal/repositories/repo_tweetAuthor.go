package repositories

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetRepo) TweetAuthor(id string) (*model.Tweets, error) {
	tweet := model.Tweets{}
	result := t.db.Model(&tweet).Select("id", "username").Where("id=?", id).Scan(&tweet)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tweet, nil
}
