package token

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Maker is an interface for managing tokens
type Token interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(ctx context.Context, userID uuid.UUID, duration time.Duration) (string, *Payload, error)
	// VerifyToken checks if the token is valid or not
	VerifyToken(ctx context.Context, token string) (*Payload, error)
}
