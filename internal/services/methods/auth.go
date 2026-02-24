package methods

import (
	database "URL-Shortener/internal/repository"
)

type AuthService struct {
	repo database.Auth
}

func NewAuthService(repo database.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
