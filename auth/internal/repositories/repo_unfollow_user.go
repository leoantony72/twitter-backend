package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"gorm.io/gorm"
)

func (u *UserPostgresRepo) UnfollowUser(follow model.User_followers) error {
	user := model.User{}
	result := u.db.Model(&follow).Where("follower = ? AND followee = ?", follow.Follower, follow.Followee).Delete(&follow)
	if result.Error != nil {
		return result.Error
	}

	//update follower -> following count -1
	u.db.Model(&user).Where("id", follow.Follower).Update("following", gorm.Expr("following - 1"))

	//update followee - followers count -1
	u.db.Model(&user).Where("id", follow.Followee).Update("followers", gorm.Expr("followers - 1"))

	return nil
}