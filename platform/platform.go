package platform

import (
	"context"
	"pharma-backend/internal/constants/model/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token interface {
	GenerateAccessToken(ctx context.Context, userID, role string, expiresAt time.Duration) (string, error)
	GenerateRefreshToken(ctx context.Context) string
	VerifyToken(signingMethod jwt.SigningMethod, token string) (bool, *dto.AccessToken, error)
}
