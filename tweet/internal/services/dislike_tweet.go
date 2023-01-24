package services

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetService) DislikeTweet(like model.Like) error {
	err := t.repo.DislikeTweet(like)
	if err != nil {
		return err
	}
	return nil
}
