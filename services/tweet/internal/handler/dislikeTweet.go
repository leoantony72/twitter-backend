package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetHandler) DislikeTweet(c *gin.Context) {
	id := c.Param("id")
	username := c.Value("username").(string)
	
	like := model.Like{TweetId: id, Username: username}
	err := t.tweet_service.DislikeTweet(like)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Disliked Tweet"})
}
