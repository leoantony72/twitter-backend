package services

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetService) GetTweetByUser(username string) (*[]model.Tweets, error) {
	tweets, err := t.repo.GetTweetByUser(username)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
