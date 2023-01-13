package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) UnfollowUser(follow model.User_followers) error {
	result := u.db.Model(&follow).Where("follower = ? AND followee = ?", follow.Follower, follow.Followee).Delete(&follow)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
