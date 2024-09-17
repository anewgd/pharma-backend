package user

import (
	"context"

	"pharma-backend/internal/constants/dbinstance"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/db"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/storage"
	"pharma-backend/platform/logger"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type user struct {
	db  dbinstance.DBInstance
	log logger.Logger
}

func Init(db dbinstance.DBInstance, log logger.Logger) storage.User {
	return &user{
		db:  db,
		log: log,
	}
}

func (u *user) Create(ctx context.Context, params dto.CreateUserRequest) (*dto.User, error) {
	user, err := u.db.CreateUser(ctx, db.CreateUserParams{
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
	})

	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not create user")
		u.log.Error(ctx, "unable to create user", zap.Error(err), zap.Any("user", params))
		return nil, err
	}
	return &dto.User{
		UserID:    user.UserID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
func (u *user) Get(ctx context.Context, username string) (*dto.User, error) {
	user, err := u.db.GetUser(ctx, username)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			err = errors.ErrNoRecordFound.New("user not found")
			u.log.Error(ctx, "unable to get user", zap.Error(err), zap.String("user-name", username))
			return nil, err
		}
		err = errors.ErrWriteError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to get user", zap.Error(err), zap.String("user-name", username))
		return nil, err
	}

	return &dto.User{
		UserID:    user.UserID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
func (u *user) CheckUserExists(ctx context.Context, email string) (bool, error) {
	count, err := u.db.UserByEmailExists(ctx, email)
	if err != nil {
		err := errors.ErrReadError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to read the user", zap.Error(err), zap.Any("user-email", email))
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
