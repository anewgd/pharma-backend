package userSession

import (
	"context"

	"pharma-backend/internal/constants/dbinstance"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/db"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/storage"
	"pharma-backend/platform/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type userSession struct {
	db  dbinstance.DBInstance
	log logger.Logger
}

func Init(db dbinstance.DBInstance, log logger.Logger) storage.UserSession {

	return &userSession{
		db:  db,
		log: log,
	}
}

func (us *userSession) Create(ctx context.Context, param dto.CreateUserSession) (*dto.UserSession, error) {
	usrSession, err := us.db.CreateUserSession(ctx, db.CreateUserSessionParams{
		UserID:       param.UserID,
		RefreshToken: param.RefreshToken,
		ExpiresAt:    param.ExpiresAt,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not create user session")
		us.log.Error(ctx, "unable to create user session", zap.Error(err), zap.Any("session", param))
		return nil, err
	}
	return &dto.UserSession{
		SessionID:    usrSession.SessionID,
		UserID:       usrSession.UserID,
		RefreshToken: usrSession.RefreshToken,
		ExpiresAt:    usrSession.ExpiresAt,
	}, nil
}
func (us *userSession) Get(ctx context.Context, sessionID uuid.UUID) (*dto.UserSession, error) {
	usrSession, err := us.db.GetUserSession(ctx, sessionID)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			err = errors.ErrUnauthorized.New("session not found")
			us.log.Error(ctx, "unable to get user session", zap.Error(err), zap.Any("session_id", sessionID))
			return nil, err
		}
		err = errors.ErrWriteError.Wrap(err, "could not retrieve user session")
		us.log.Error(ctx, "unable to retrieve user session", zap.Error(err), zap.Any("session_id", sessionID))
		return nil, err
	}

	return &dto.UserSession{
		SessionID:    usrSession.SessionID,
		UserID:       usrSession.UserID,
		RefreshToken: usrSession.RefreshToken,
		IsBlocked:    usrSession.IsBlocked,
		ExpiresAt:    usrSession.ExpiresAt,
		CreatedAt:    usrSession.CreatedAt,
	}, nil
}
func (us *userSession) Delete(ctx context.Context, userID uuid.UUID) error {
	err := us.db.DeleteUserSession(ctx, userID)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete user sessions")
		us.log.Error(ctx, "unable to delete user sessions", zap.Error(err), zap.Any("user_id", userID))
		return err
	}
	return nil
}
func (us *userSession) DeleteAll(ctx context.Context) error {
	err := us.db.DeleteAllUserSessions(ctx)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete sessions")
		us.log.Error(ctx, "unable to delete sessions", zap.Error(err))
		return err
	}
	return nil
}
