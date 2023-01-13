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
	// fmt.Println(user)
	// userData, _ := json.Marshal(user)
	// fmt.Println(string(userData))
	c.JSON(200, user)
}
