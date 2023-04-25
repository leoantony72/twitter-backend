package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

func (u *UserHandler) CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}

	err = u.userUseCase.Create(user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please check your values"})
		return
	}
	c.JSON(201, gin.H{"message": "User Created Successfully"})
}
