package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

var duplicateFollow string = "ERROR: duplicate key value violates unique constraint \"user_followers_pkey\" (SQLSTATE 23505)"

func (u *UserHandler) Follow(c *gin.Context) {
	follow := model.User_followers{}
	follower := c.Value("id").(string)
	followee := c.Query("follow")

	follow.Follower = follower
	follow.Followee = followee

	fmt.Printf("follower : %v and followee: %s", follower, followee)
	err := u.userUseCase.FollowUser(follow)
	if err != nil {
		if err.Error() == duplicateFollow {
			c.JSON(400, gin.H{"message": "Alredy Followed"})
			return
		}
		c.JSON(400, gin.H{"message": "Failed to follow", "err": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Followed User"})
}
