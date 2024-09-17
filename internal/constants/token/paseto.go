package token

import (
	"context"
	"time"

	"pharma-backend/internal/constants/errors"
	"pharma-backend/platform/logger"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"go.uber.org/zap"
)

type Paseto struct {
	paseto       *paseto.V2
	symmerticKey []byte
	log          logger.Logger
}

func InitPaseto(symmerticKey string, log logger.Logger) Token {

	return &Paseto{
		paseto:       paseto.NewV2(),
		symmerticKey: []byte(symmerticKey),
		log:          log,
	}
}

func (p *Paseto) CreateToken(ctx context.Context, userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	paylood, err := NewPayload(userID, duration)
	if err != nil {
		p.log.Error(ctx, "unable to create token", zap.Error(err))
		return "", nil, err
	}
	token, err := p.paseto.Encrypt(p.symmerticKey, paylood, nil)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "unable to encrypt token payload")
		p.log.Error(ctx, "unable to create token", zap.Error(err))
		return "", nil, err
	}
	return token, paylood, err
}

// VerifyToken checks if the token is valid or not
func (p *Paseto) VerifyToken(ctx context.Context, token string) (*Payload, error) {
	payload := &Payload{}
	err := p.paseto.Decrypt(token, p.symmerticKey, payload, nil)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "unable to decrypt token payload")
		p.log.Error(ctx, "unable to verify token", zap.Error(err))
		return nil, err
	}

	err = payload.Validate()
	if err != nil {
		p.log.Error(ctx, "unable to verify token", zap.Error(err))
		return nil, err
	}
	return payload, nil
}
