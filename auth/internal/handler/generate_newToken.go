package handler

import (
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) GenerateNewToken(c *gin.Context) {
	refresh_Token, err := c.Cookie("refresh-Token")
	// client_username := c.Value("username")
	if err != nil {
		c.JSON(400, gin.H{"message": "Please Login"})
		return
	}
	_, claims, err := u.userUseCase.GetTokenClaims(refresh_Token)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please Login Again"})
		return
	}
	user, err := u.userUseCase.GetToken(refresh_Token)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please Try Again"})
		return
	}
	if user.Username != claims["name"] && user.Token != refresh_Token {
		c.JSON(400, gin.H{"message": "Invalid token please login again"})
		return
	}
	accessToken, err := u.userUseCase.GenerateAccessToken(user.Username, user.Id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please Try Again"})
		return
	}
	c.SetCookie("access-Token", accessToken, 3600, "/", "", false, true)
	c.JSON(201, gin.H{"message": "Generated New Acess-Token"})
}
