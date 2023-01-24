package middleware

import "github.com/gin-gonic/gin"

func (t *TweetMiddleware) CheckAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		client_username := c.Value("username").(string)
		if id == "" {
			c.JSON(400, gin.H{"message": "provide the Tweet Id"})
			c.Abort()
			return
		}
		tweet, err := t.tweetService.TweetAuthor(id)
		if err != nil {
			c.JSON(500, gin.H{"message": "something went wrong"})
			c.Abort()
			return
		}
		if tweet.Username != client_username {
			c.JSON(400, gin.H{"message": "Invalid Operation"})
			c.Abort()
			return
		}
		c.Next()

	}
}
