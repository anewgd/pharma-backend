package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AccessToken struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	// AccessToken is the access token for the current login
	AccessToken string `form:"access_token" query:"access_token" json:"access_token,omitempty"`
	// IDToken is the OpenID specific JWT token
	RefreshToken string `form:"refresh_token" query:"refresh_token" json:"refresh_token,omitempty"`
	// TokenType is the type of token
	TokenType string `form:"token_type" query:"token_type" json:"token_type,omitempty"`
	// ExpiresAt is time the access token is going to be expired.
	ExpiresAt time.Time `json:"expires_at"`
}
