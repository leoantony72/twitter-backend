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
		tweet.GET("/id/:id", handler.GetTweetById) //get tweet by id
		tweet.POST("/create", m.Authorization(), handler.CreateTweet)
		tweet.DELETE("id/:id/delete", m.Authorization(), m.CheckAuthor(), handler.DeleteTweet)
		tweet.POST("id/:id/like", m.Authorization(), handler.LikeTweet)
		tweet.POST("id/:id/dislike", m.Authorization(), handler.DislikeTweet)
		tweet.POST("id/:id/retweet", m.Authorization(), handler.ReTweet)
		tweet.DELETE("id/:id/retweet", m.Authorization(), handler.DeleteReTweet)
		// tweet.GET("/username/:username", handler.GetTweetByUsername)
		//@Depricated 
		//get tweet by username. returns array of tweets
	}

	return handler
}
