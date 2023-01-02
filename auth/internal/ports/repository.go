package ports

import "github.com/leoantony72/twitter-backend/auth/internal/model"

type UserRepository interface {
	Create(user model.User) error
	Update(user model.User) error
	Delete(id string) error
	GetUserbyId(id string) (*model.User, error)
	Login(username string) (*model.User, error)
}
