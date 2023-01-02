package repositories

import (
	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"github.com/leoantony72/twitter-backend/auth/internal/ports"
	"gorm.io/gorm"
)

type UserPostgresRepo struct {
	db *gorm.DB
}

func NewUserPostgresRepo(db *gorm.DB) ports.UserRepository {
	return &UserPostgresRepo{
		db: db,
	}
}

func (u *UserPostgresRepo) Create(user model.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *UserPostgresRepo) Update(user model.User) error {
	result := u.db.UpdateColumns(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *UserPostgresRepo) Delete(id string) error {
	result := u.db.Delete(id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *UserPostgresRepo) GetUserbyId(id string) (*model.User, error) {
	user := model.User{}
	result := u.db.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (u *UserPostgresRepo) Login(username string) (*model.User, error) {
	user := model.User{}
	result := u.db.Find(&user, username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
