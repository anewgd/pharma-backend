package token

import (
	"net/http"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/anewgd/pharma_backend/util"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmerticKey []byte
}

func NewPasetoMaker(symmerticKey string) (Maker, error) {
	if len(symmerticKey) != chacha20poly1305.KeySize {
		return nil, util.AuthenticationError.New("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmerticKey: []byte(symmerticKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(userID uuid.UUID, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, role, duration)
	if err != nil {
		return "", nil, err
	}
	token, err := maker.paseto.Encrypt(maker.symmerticKey, payload, nil)
	if err != nil {
		return "", nil, err
	}
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmerticKey, payload, nil)

	if err != nil {
		return nil, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to decrypt authentication payload"), http.StatusInternalServerError, "internal error")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
