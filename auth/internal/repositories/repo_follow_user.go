package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"gorm.io/gorm"
)

func (u *UserPostgresRepo) FollowUser(follow model.User_followers) error {
	user := model.User{}
	err := u.db.Model(&follow).Create(&follow)
	//update follower -> following count +1
	u.db.Model(&user).Where("id", follow.Follower).Update("following", gorm.Expr("following + 1"))

	//update followee - followers count +1
	u.db.Model(&user).Where("id", follow.Followee).Update("followers", gorm.Expr("followers + 1"))

	return err.Error
}
