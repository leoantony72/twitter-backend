package handler

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) GetById(c *gin.Context) {
	id := c.Query("id")

	user, err := u.userUseCase.GetUserbyId(id)
	if err != nil {
		c.JSON(200, gin.H{"message": "user not found"})
	}
	c.JSON(200, user)
}
