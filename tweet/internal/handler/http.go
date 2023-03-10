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
	//service health endpoint
	r.GET("/check", handler.Check)
	tweet := r.Group("/tweet")
	{
		tweet.GET("/id/:id", handler.GetTweetById)                   //get tweet by id
		tweet.GET("/username/:username", handler.GetTweetByUsername) //get tweet by username. returns array of tweets
		tweet.POST("/create", m.Authorization(), handler.CreateTweet)
		tweet.DELETE("/:id/delete", m.Authorization(), m.CheckAuthor(), handler.DeleteTweet)
		tweet.POST("/:id/like", m.Authorization(), handler.LikeTweet)
		tweet.POST("/:id/dislike", m.Authorization(), handler.DislikeTweet)
		tweet.POST("/:id/retweet", m.Authorization(), handler.ReTweet)
		tweet.DELETE("/:id/retweet", m.Authorization(), handler.DeleteReTweet)
	}

	return handler
}
