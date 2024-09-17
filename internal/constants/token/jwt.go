package token

import (
	"context"
	"errors"
	"time"

	errs "pharma-backend/internal/constants/errors"
	"pharma-backend/platform/logger"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var ErrExpiredToken = errors.New("token has expired")

// JWTMaker is a JWT maker
type JWT struct {
	secretKey string
	log       logger.Logger
}

// NewJWTMaker creates a new JWTMaker
func InitJWT(secretKey string, log logger.Logger) Token {
	return &JWT{secretKey: secretKey, log: log}
}

func (j *JWT) CreateToken(ctx context.Context, userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		j.log.Error(ctx, "unable to create token", zap.Error(err))
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(j.secretKey))
	if err != nil {
		err = errs.ErrWriteError.Wrap(err, "unable to sign token")
		j.log.Error(ctx, "unable to create token", zap.Error(err))
	}
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (j *JWT) VerifyToken(ctx context.Context, token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			err := errs.ErrUnauthorized.New("invalid signing method")
			j.log.Error(ctx, "unable to verify token", zap.Error(err))
			return nil, err
		}
		return []byte(j.secretKey), nil

	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		err = errs.ErrUnauthorized.Wrap(err, "invalid token")
		j.log.Error(ctx, "unable to verify token", zap.Error(err))
		return nil, err
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		err = errs.ErrWriteError.New("failed to cast token claims to payload")
		j.log.Error(ctx, "unable to verify token", zap.Error(err))
		return nil, err
	}
	return payload, nil
}
