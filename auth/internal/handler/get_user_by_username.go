package handler

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := u.userUseCase.GetUserbyUsername(username)
	if err != nil {
		c.JSON(200, gin.H{"message": "user not found"})
	}
	c.JSON(200, user)
}
