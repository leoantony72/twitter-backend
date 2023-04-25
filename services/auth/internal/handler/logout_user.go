package handler

import "github.com/gin-gonic/gin"

func (u *UserHandler) LogoutUser(c *gin.Context) {
	id := c.Value("id").(string)

	err := u.userUseCase.Logout(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.SetCookie("access-Token", "", -1, "/", "", false, true)
	c.SetCookie("refresh-Token", "", -1, "/", "", false, true)
	c.JSON(201, gin.H{"message": "Logout Success"})
}
