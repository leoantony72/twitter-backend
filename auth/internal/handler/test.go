package handler

import (
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Test(c *gin.Context) {
	username := c.Value("username")
	id := c.Value("id")
	c.JSON(200, gin.H{"id": id, "username": username})
}
