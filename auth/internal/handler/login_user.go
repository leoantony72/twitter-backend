package handler

import "github.com/gin-gonic/gin"

type LoginCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	credentials := LoginCredential{}

	err := c.BindJSON(&credentials)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = u.userUseCase.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	//also send JWT token to the user
	c.JSON(200, gin.H{"message": "Login Success"})
}
