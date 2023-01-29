package repositories

import "github.com/leoantony72/twitter-backend/auth/internal/model"

func (u *UserPostgresRepo) Create(user model.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	redis_key := "users:" + user.Username + ":" + "metadata"
	err := u.redis.HSet(
		ctx, redis_key,
		"id", user.Id,
		"username", user.Username,
		"email", user.Email,
		"date_created", user.Date_created,
		"following", 0,
		"followers", 0,
	)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}
