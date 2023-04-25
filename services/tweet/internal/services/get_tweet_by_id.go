package services

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetService) GetTweetById(id string) (*model.Tweets, error) {
	tweet, err := t.repo.GetTweetById(id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
