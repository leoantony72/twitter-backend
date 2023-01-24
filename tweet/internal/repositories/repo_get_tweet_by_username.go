package repositories

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetRepo) GetTweetByUser(username string) (*[]model.Tweets, error) {
	tweet := []model.Tweets{}
	result := t.db.Model(&tweet).Select("id", "username", "created_at", "like_count", "retweet_count").Where("username=?", username).Scan(&tweet)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tweet, nil
}
