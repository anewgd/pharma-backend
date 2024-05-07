package util

import (
	"context"
	"net/http"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	userIDPayload := ctx.Value(UserID)

	if userIDPayload == nil {
		return uuid.UUID{}, NewErrorResponse(errorx.InternalError.New("no value found in context with key %s", string(UserID)), http.StatusInternalServerError, "internal error")
	}
	userID, ok := (userIDPayload).(uuid.UUID)
	if !ok {
		return uuid.UUID{}, NewErrorResponse(errorx.InternalError.New("cannot extract user id from context"), http.StatusInternalServerError, "internal error")
	}

	return userID, nil
}

func GetUserRole(ctx context.Context) (string, error) {
	rolePayload := ctx.Value(Role)
	if rolePayload == nil {
		return "", NewErrorResponse(errorx.InternalError.New("no value found in context with key %s", string(Role)), http.StatusInternalServerError, "internal error")
	}
	userRole, ok := (rolePayload).(string)
	if !ok {
		return "", NewErrorResponse(errorx.InternalError.New("failed to extract user role from context"), http.StatusInternalServerError, "internal error")
	}
	return userRole, nil
}
