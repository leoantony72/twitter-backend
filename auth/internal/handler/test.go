package handler

import (
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Test(c *gin.Context) {
	// userData, _ := json.Marshal(user)
	// fmt.Println(string(userData))
	c.JSON(200, gin.H{"message": "you are in ADMIN"})
}
