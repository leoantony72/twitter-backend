package services

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetService) ReTweet(retweet model.Retweet) error {
	err := t.repo.ReTweet(retweet)
	if err != nil {
		return err
	}
	return nil
}
