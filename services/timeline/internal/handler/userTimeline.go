package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (t *TimelineHandler) UserTimeline(c *gin.Context) {
	user := c.Query("username")
	Tstart := c.Query("start")
	start, err := strconv.Atoi(Tstart)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	tweets, err := t.timeline_service.GetUserTimeline(user, start)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": tweets})
}
