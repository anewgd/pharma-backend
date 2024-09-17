package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/platform"
	"pharma-backend/platform/logger"
	"pharma-backend/platform/util"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type JWT struct {
	logger     logger.Logger
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func JWTInit(logger logger.Logger, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) platform.Token {
	return &JWT{
		logger:     logger,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}
func (j *JWT) GenerateAccessToken(ctx context.Context, userID, role string, expiresAt time.Duration) (string, error) {
	claims := dto.AccessToken{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
			Issuer:    "pharma",
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodPS512, claims).SignedString(j.privateKey)
	if err != nil {
		j.logger.Error(ctx, "could not generate access token", zap.Error(err))
		return "", errors.ErrWriteError.Wrap(err, "could not generate access token")
	}

	return token, err
}
func (j *JWT) GenerateRefreshToken(ctx context.Context) string {
	return util.GenerateRandomString(32)
}
func (j *JWT) VerifyToken(signingMethod jwt.SigningMethod, token string) (bool, *dto.AccessToken, error) {
	claims := &dto.AccessToken{}

	segments := strings.Split(token, ".")
	if len(segments) < 3 {
		return false, claims, errors.ErrInvalidToken.New("malformed access token")
	}

	//TODO:
	// err := signingMethod.Verify(token, segments[2], j.publicKey)
	// if err != nil {
	// 	return false, claims, errors.ErrAuthError.Wrap(err, "token verification failed")
	// }

	if _, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSAPSS); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return j.publicKey, nil
	}); err != nil {
		return false, claims, errors.ErrAuthError.WrapWithNoMessage(err)
	}
	return true, claims, nil
}
