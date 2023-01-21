package middleware

import "github.com/leoantony72/twitter-backend/tweet/internal/ports"

type TweetMiddleware struct {
	tweetService ports.TweetService
}

func NewTweetMiddleware(s ports.TweetService) *TweetMiddleware {
	return &TweetMiddleware{
		tweetService: s,
	}
}
