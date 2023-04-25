package handler

import "github.com/gin-gonic/gin"

func (u *UserHandler) DeleteUser(c *gin.Context) {

	username := c.Value("username").(string)
	// username := c.Query("username")

	err := u.userUseCase.Delete(username)
	if err != nil {
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(200, gin.H{"message": "User Deleted"})
}
