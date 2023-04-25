package handler

import (
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := u.userUseCase.GetUserbyUsername(username)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, user)
}
