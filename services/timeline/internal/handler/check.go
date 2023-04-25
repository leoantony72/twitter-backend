package handler

import "github.com/gin-gonic/gin"

func (t *TimelineHandler) Check(c *gin.Context) {
	c.Status(200)
}