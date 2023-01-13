package handler

import "github.com/gin-gonic/gin"

func (u *UserHandler) DeleteUser(c *gin.Context) {

	id := c.Value("id").(string)
	// id := c.Query("id")

	err := u.userUseCase.Delete(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(200, gin.H{"message": "User Deleted"})
}
