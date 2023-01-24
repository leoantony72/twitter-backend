package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/internal/middleware"
	"github.com/leoantony72/twitter-backend/tweet/internal/ports"
)

type TweetHandler struct {
	tweet_service ports.TweetService
}

func NewTweetHandler(s ports.TweetService, m *middleware.TweetMiddleware, r *gin.Engine) *TweetHandler {
	handler := &TweetHandler{
		tweet_service: s,
	}
	tweet := r.Group("/tweet")
	{
		tweet.GET("/id/:id", handler.GetTweetById)                   //get tweet by id
		tweet.GET("/username/:username", handler.GetTweetByUsername) //get tweet by username. returns array of tweets
		tweet.POST("/create", m.Authorization(), handler.CreateTweet)
		// tweet.DELETE("/:id/Delete")
		// tweet.POST("/:id/like")
		// tweet.DELETE("/:id/dislike")
		// tweet.POST("/:id/retweet")
		// tweet.DELETE("/:id/retweet")
	}

	return handler
}
