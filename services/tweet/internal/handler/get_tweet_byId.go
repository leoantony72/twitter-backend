package handler

import (
	"github.com/gin-gonic/gin"
)

func (t *TweetHandler) GetTweetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"message": "Tweet Id not specified"})
		return
	}
	tweet, err := t.tweet_service.GetTweetById(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// jsonData, _ := json.Marshal(&tweet)
	c.JSON(200, tweet)
}
