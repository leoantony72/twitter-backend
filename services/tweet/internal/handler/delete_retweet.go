package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetHandler) DeleteReTweet(c *gin.Context) {
	id := c.Param("id")
	username := c.Value("username").(string)

	retweet := model.Retweet{TweetId: id, Username: username}
	err := t.tweet_service.DeleteReTweet(retweet)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "retweet deleted"})
}
