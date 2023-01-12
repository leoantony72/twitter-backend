package handler

import (
	"github.com/gin-gonic/gin"
)

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

	accessToken, refreshToken, err := u.userUseCase.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// stringT := utils.ValidateJwt(accessToken)
	// fmt.Printf("%v\n", stringT)

	err = u.userUseCase.AddToken(credentials.Username, refreshToken)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	//also send JWT token to the user
	c.SetCookie("access-Token", accessToken, 3600, "/", "", false, true)
	c.SetCookie("refresh-Token", refreshToken, 36000, "/", "", false, true)
	c.JSON(200, gin.H{"message": "Login Success"})
}
