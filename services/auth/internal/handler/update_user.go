package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

func (u *UserHandler) UpdateUser(c *gin.Context) {
	user := model.User{}
	id := c.Value("id").(string)
	prev_username := c.Value("username").(string)

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}
	user.Id = id
	err = u.userUseCase.Update(user,prev_username)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please check your values", "err": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User Updated Succesfully"})
}
