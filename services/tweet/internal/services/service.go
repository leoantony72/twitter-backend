package services

import "github.com/leoantony72/twitter-backend/tweet/internal/ports"

type TweetService struct {
	repo ports.TweetRepository
}

func NewTweetService(repo ports.TweetRepository) ports.TweetService {
	return &TweetService{
		repo: repo,
	}
}
