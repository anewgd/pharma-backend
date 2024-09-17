package token

import (
	"time"

	"pharma-backend/internal/constants/errors"

	"github.com/golang-jwt/jwt/v4"

	"github.com/google/uuid"
)

// Payload contains the payload data for the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(userID uuid.UUID, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrWriteError.Wrap(err, "failed to create a access token ID")
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Validate() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.ErrUnauthorized.New("expired token")
	}
	return nil
}
