package handler

import "github.com/gin-gonic/gin"

func (t *TweetHandler) Check(c *gin.Context) {
	c.Status(200)
}
