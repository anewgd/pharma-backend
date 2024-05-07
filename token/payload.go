package token

import (
	"errors"
	"net/http"
	"time"

	"github.com/anewgd/pharma_backend/util"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("invalid token")

// Payload contains the payload data for the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(userID uuid.UUID, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, util.NewErrorResponse(errorx.InternalError.Wrap(err, "cannot generate token id"), http.StatusInternalServerError, "internal error")
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return util.NewErrorResponse(util.AuthorizationError.New("token has expired"), http.StatusUnauthorized, "token has expired")
	}
	return nil
}
