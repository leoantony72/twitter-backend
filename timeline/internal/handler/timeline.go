package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (t *TimelineHandler) Timeline(c *gin.Context) {
	user := c.Value("username").(string)
	Tstart := c.Query("start")
	start, err := strconv.Atoi(Tstart)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	fmt.Println(start, user)
	tweets, err := t.timeline_service.GetTimeline(user, start)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": tweets})
}
