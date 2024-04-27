package service

import (
	"time"

	db "github.com/anewgd/pharma_backend/data/sqlc"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func newCreateUserResponse(user db.User) CreateUserResponse {
	return CreateUserResponse{
		UserID:   user.UserID.String(),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Username              string    `json:"username"`
	Email                 string    `json:"email"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}
