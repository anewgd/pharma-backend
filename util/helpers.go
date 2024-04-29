package util

import (
	"context"
	"fmt"

	"github.com/anewgd/pharma_backend/util/token"
	"github.com/gin-gonic/gin"
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

func GetContextWithValues(ctx *gin.Context) (context.Context, error) {

	c := ctx.Request.Context()
	payload, ok := ctx.Get(AuthorizationPayloadKey)
	if !ok {
		return nil, fmt.Errorf("cannot find authorization payload")
	}

	usrPayload, ok := (payload).(*token.Payload)
	if !ok {
		return nil, fmt.Errorf("can't get user id")
	}

	c = context.WithValue(c, UserID, usrPayload.UserID)
	c = context.WithValue(c, Role, usrPayload.Role)

	return c, nil
}
