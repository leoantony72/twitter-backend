package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"gorm.io/gorm"
)

func (u *UserPostgresRepo) FollowUser(follow model.User_followers) error {
	user := model.User{}
	err := u.db.Model(&follow).Create(&follow)
	if err.Error != nil {
		return err.Error
	}
	//update follower -> following count +1
	u.db.Model(&user).Where("username", follow.Follower).Update("following_count", gorm.Expr("following_count + 1"))

	//update followee - followers count +1
	u.db.Model(&user).Where("username", follow.Followee).Update("followers_count", gorm.Expr("followers_count + 1"))

	return err.Error
}
