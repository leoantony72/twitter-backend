package services

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetService) DeleteReTweet(retweet model.Retweet) error {
	err := t.repo.DeleteReTweet(retweet)
	if err != nil {
		return err
	}
	return nil
}
