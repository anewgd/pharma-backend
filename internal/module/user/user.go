package user

import (
	"context"
	"time"

	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/module"
	"pharma-backend/internal/storage"
	"pharma-backend/platform"
	"pharma-backend/platform/logger"
	"pharma-backend/platform/util"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type user struct {
	userPersistent        storage.User
	userSessionPersistent storage.UserSession
	token                 platform.Token
	log                   logger.Logger
}

func Init(log logger.Logger, userPersistent storage.User, token platform.Token, userSessionPersistent storage.UserSession) module.User {
	return &user{
		userPersistent:        userPersistent,
		userSessionPersistent: userSessionPersistent,
		token:                 token,
		log:                   log,
	}
}

func (u *user) CreateUser(ctx context.Context, param dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	exists, err := u.userPersistent.CheckUserExists(ctx, param.Email)
	if err != nil {
		return nil, err
	} else if exists {
		err = errors.ErrDataExists.New("user with this email already exists")
		u.log.Error(ctx, "duplicated data", zap.String("user-email", param.Email))
		return nil, err
	}

	hashedPassword, err := util.HashAndSalt(ctx, []byte(param.Password), u.log)
	if err != nil {
		return nil, err
	}

	user, err := u.userPersistent.Create(ctx, dto.CreateUserRequest{
		Username: param.Username,
		Password: string(hashedPassword),
		Email:    param.Email,
	})

	if err != nil {
		return nil, err
	}

	newUser := dto.CreateUserResponse{
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
	return &newUser, nil
}
func (u *user) GetUser(ctx context.Context, param dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	user, err := u.userPersistent.Get(ctx, param.Username)
	if err != nil {
		return nil, err
	}

	isCorrectPassword := util.CompareHashAndPassword(user.Password, param.Password)
	if !isCorrectPassword {
		err = errors.ErrUnauthorized.New("incorrect password")
		u.log.Error(ctx, "unable to get user", zap.Error(err))
		return nil, err
	}

	// accessToken, accessTokenPayload, err := u.token.CreateToken(ctx, user.UserID, viper.GetDuration("ACCESS_TOKEN_DURATION"))

	accessToken, err := u.token.GenerateAccessToken(ctx, user.UserID.String(), user.Role, viper.GetDuration("token.access_token_duration"))
	if err != nil {
		return nil, err
	}

	refreshToken := u.token.GenerateRefreshToken(ctx)

	// refreshToken, refreshTokenPayload, err := u.token.CreateToken(ctx, user.UserID, viper.GetDuration("REFRESH_TOKEN_DURATION"))
	// if err != nil {
	// 	return nil, err
	// }

	err = u.userSessionPersistent.Delete(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	refreshTokenExpiresAt := time.Now().Add(viper.GetDuration("token.refresh_token_duration"))
	//TODO: ADD Token related code
	u.userSessionPersistent.Create(ctx, dto.CreateUserSession{
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenExpiresAt,
	})

	return &dto.LoginUserResponse{
		Username:              user.Username,
		Email:                 user.Email,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  time.Now().Add(viper.GetDuration("token.access_token_duration")),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
	}, nil
}
