package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

var duplicate_Error_retweet string = "ERROR: duplicate key value violates unique constraint \"retweets_pkey\" (SQLSTATE 23505)"

func (t TweetHandler) ReTweet(c *gin.Context) {
	id := c.Param("id")
	username := c.Value("username").(string)

	retweet := model.Retweet{TweetId: id, Username: username}
	err := t.tweet_service.ReTweet(retweet)
	if err != nil {
		if err.Error() == duplicate_Error_retweet {
			c.JSON(400, gin.H{"message": "Alredy retweeted"})
			return
		}
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "retweeted"})
}
