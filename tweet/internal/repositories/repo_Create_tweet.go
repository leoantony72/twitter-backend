package repositories

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetRepo) CreateTweet(tweet model.Tweets) error {

	result := t.db.Create(&tweet)
	// CheckErr(result.Error)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
