package pharmacySession

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

type pharmacistSesssion struct {
	db  dbinstance.DBInstance
	log logger.Logger
}

func Init(db dbinstance.DBInstance, log logger.Logger) storage.PharmacySession {

	return &pharmacistSesssion{
		db:  db,
		log: log,
	}
}
func (ps *pharmacistSesssion) Create(ctx context.Context, param dto.CreatePharmacistSession) (*dto.PharmacistSession, error) {
	pharmSession, err := ps.db.CreatePharmacistSession(ctx, db.CreatePharmacistSessionParams{
		SessionID:    param.SessionID,
		PharmacistID: param.PharmacistID,
		RefreshToken: param.RefreshToken,
		ExpiresAt:    param.ExpiresAt,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not create pharmacist session")
		ps.log.Error(ctx, "unable to create pharmacist session", zap.Error(err), zap.Any("session", param))
		return nil, err
	}
	return &dto.PharmacistSession{
		SessionID:    pharmSession.SessionID,
		PharmacistID: pharmSession.PharmacistID,
		RefreshToken: pharmSession.RefreshToken,
		ExpiresAt:    pharmSession.ExpiresAt,
	}, nil
}
func (ps *pharmacistSesssion) Get(ctx context.Context, sessionID uuid.UUID) (*dto.PharmacistSession, error) {
	pharmSession, err := ps.db.GetUserSession(ctx, sessionID)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			err = errors.ErrUnauthorized.New("session not found")
			ps.log.Error(ctx, "unable to get pharmacist session", zap.Error(err), zap.Any("session_id", sessionID))
			return nil, err
		}
		err = errors.ErrWriteError.Wrap(err, "could not retrieve pharmacist session")
		ps.log.Error(ctx, "unable to retrieve pharmacist session", zap.Error(err), zap.Any("session_id", sessionID))
		return nil, err
	}

	return &dto.PharmacistSession{
		SessionID:    pharmSession.SessionID,
		PharmacistID: pharmSession.UserID,
		RefreshToken: pharmSession.RefreshToken,
		IsBlocked:    pharmSession.IsBlocked,
		ExpiresAt:    pharmSession.ExpiresAt,
		CreatedAt:    pharmSession.CreatedAt,
	}, nil
}
func (ps *pharmacistSesssion) Delete(ctx context.Context, userID uuid.UUID) error {
	err := ps.db.DeletePharmacistSession(ctx, userID)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete pharmacist sessions")
		ps.log.Error(ctx, "unable to delete pharmacist sessions", zap.Error(err), zap.Any("pharmacist_id", userID))
		return err
	}
	return nil
}
func (ps *pharmacistSesssion) DeleteAll(ctx context.Context) error {
	err := ps.db.DeleteAllUserSessions(ctx)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not delete sessions")
		ps.log.Error(ctx, "unable to delete sessions", zap.Error(err))
		return err
	}
	return nil
}
