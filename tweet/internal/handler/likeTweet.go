package handler

import (
	"github.com/gin-gonic/gin"
)

var duplicate_Error string = "ERROR: duplicate key value violates unique constraint \"likes_pkey\" (SQLSTATE 23505)"

func (t *TweetHandler) LikeTweet(c *gin.Context) {
	id := c.Param("id")
	username := c.Value("username").(string)
	err := t.tweet_service.LikeTweet(id, username)
	if err != nil {
		if err.Error() == duplicate_Error {
			c.JSON(400, gin.H{"message": "Alredy Liked"})
			return
		}
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Tweet liked"})
}
