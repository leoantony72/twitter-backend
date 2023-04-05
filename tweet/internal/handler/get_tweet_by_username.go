//@@ DEPRECATED
//@ Moved to Timeline Service In "/timeline/user"


package handler


import (
	"github.com/gin-gonic/gin"
)

func (t *TweetHandler) GetTweetByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(400, gin.H{"message": "Username not provided"})
		return
	}
	tweets, err := t.tweet_service.GetTweetByUser(username)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": tweets})
}
