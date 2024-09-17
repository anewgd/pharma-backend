package util

import (
	"context"
	"pharma-backend/internal/constants"
	"pharma-backend/internal/constants/errors"

	"github.com/google/uuid"
)

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	userIdValue := ctx.Value(constants.Context("x-user-id"))
	if userIdValue == nil {
		err := errors.ErrReadError.New("failed to read user id")
		return uuid.UUID{}, err
	}

	userIDStr, ok := (userIdValue).(string)
	if !ok {
		err := errors.ErrReadError.New("failed to process user id")
		return uuid.UUID{}, err
	}

	uID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.UUID{}, errors.ErrReadError.New("failed to parse user id")
	}

	return uID, nil
}

func GetUserRole(ctx context.Context) (string, error) {
	role := ctx.Value(constants.Context("x-user-role"))
	if role == nil {
		err := errors.ErrReadError.New("user role was not found in context")
		return "", err
	}
	userRole, ok := (role).(string)
	if !ok {
		return "", errors.ErrReadError.New("failed to extract user role")
	}
	return userRole, nil
}
