package services

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"github.com/leoantony72/twitter-backend/tweet/internal/utils"
)

func (t *TweetService) CreateTweet(tweet model.Tweets) error {
	id := utils.GenerateID()
	tweet.Id = id.String()
	err := t.repo.CreateTweet(tweet)
	if err != nil {
		return err
	}
	return nil
}
