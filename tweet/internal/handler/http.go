package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/internal/ports"
)

type TweetHandler struct {
	tweet_service ports.TweetService
}

func NewTweetHandler(s ports.TweetService, r *gin.Engine) *TweetHandler {
	handler := &TweetHandler{
		tweet_service: s,
	}
	tweet := r.Group("/tweet")
	{
		tweet.GET("/id/:id")             //get tweet by id
		tweet.GET("/username/:username") //get tweet by username. returns array of tweets
		tweet.POST("/Create")
		tweet.DELETE("/:id/Delete")
		tweet.POST("/:id/like")
		tweet.DELETE("/:id/dislike")
		tweet.POST("/:id/retweet")
		tweet.DELETE("/:id/retweet")
	}

	return handler
}
