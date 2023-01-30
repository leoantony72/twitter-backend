package ports

import (
	"github.com/golang-jwt/jwt"
	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

type UserUseCase interface {
	Create(user model.User) error
	Update(user model.User) error
	Delete(id string) error

	Login(username string, password string) (string, string, error)
	Logout(id string) error
	AddToken(username, token string) error
	GetToken(refreshToken string) (model.User, error)
	GetTokenClaims(token string) (*jwt.Token, jwt.MapClaims, error)
	GenerateAccessToken(name, Id string) (string, error)

	GetUserbyUsername(id string) (*model.User, error)
	FollowUser(follow model.User_followers) error
	UnfollowUser(follow model.User_followers) error
}
