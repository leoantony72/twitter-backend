package repositories

import (
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
