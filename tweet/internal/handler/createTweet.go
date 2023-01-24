package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetHandler) CreateTweet(c *gin.Context) {
	tweet := model.Tweets{}
	err := c.BindJSON(&tweet)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	tweet.Username = c.Value("username").(string)
	err = t.tweet_service.CreateTweet(tweet)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Tweet Created"})
}
