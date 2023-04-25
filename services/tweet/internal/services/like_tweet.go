package services

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

func (t *TweetService) LikeTweet(like model.Like) error {
	err := t.repo.LikeTweet(like)
	if err != nil {
		return err
	}
	return nil
}
