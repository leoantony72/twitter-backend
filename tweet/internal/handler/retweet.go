package handler

import "github.com/gin-gonic/gin"

func (t TweetHandler) ReTweet(c *gin.Context) {
	id := c.Param("id")
	username := c.Value("username").(string)

	err := t.tweet_service.ReTweet(id, username)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}
	c.JSON(201, gin.H{"message": "retweeted"})
}
