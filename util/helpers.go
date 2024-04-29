package util

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	userIDPayload := ctx.Value(UserID)

	if userIDPayload == nil {
		return uuid.UUID{}, fmt.Errorf("no value found in context")
	}
	userID, ok := (userIDPayload).(uuid.UUID)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("cannot find user id")
	}

	return userID, nil
}

func GetUserRole(ctx context.Context) (string, error) {
	rolePayload := ctx.Value(Role)
	if rolePayload == nil {
		return "", fmt.Errorf("no value found in context")
	}
	userRole, ok := (rolePayload).(string)
	if !ok {
		return "", fmt.Errorf("cannot find user role id")
	}
	return userRole, nil
}
