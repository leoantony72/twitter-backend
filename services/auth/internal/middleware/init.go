package middleware

import "github.com/leoantony72/twitter-backend/auth/internal/ports"

type UserMiddleware struct {
	userUseCase ports.UserUseCase
}

func NewUserMiddleware(u ports.UserUseCase) *UserMiddleware {
	handler := &UserMiddleware{userUseCase: u}
	return handler
}
