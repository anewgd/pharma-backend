package service

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/token"
	"github.com/anewgd/pharma_backend/util"
	"github.com/jackc/pgx/v5"
	"github.com/joomcode/errorx"
)

type UserService interface {
	CreateUser(ctx context.Context, userReq CreateUserRequest) (CreateUserResponse, error)
	LoginUser(ctx context.Context, userReq LoginUserRequest) (LoginUserResponse, error)
	GetTokenMaker() *token.Maker
}

type UserServ struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewUserService(store db.Store) (*UserServ, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		return nil, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to load configuration file"), http.StatusInternalServerError, "internal server error")
	}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create token maker"), http.StatusInternalServerError, "internal error")
	}
	return &UserServ{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}, nil
}

func (u *UserServ) GetTokenMaker() *token.Maker {
	return &u.tokenMaker
}

func (u *UserServ) CreateUser(ctx context.Context, userReq CreateUserRequest) (CreateUserResponse, error) {
	usrResp := CreateUserResponse{}

	if err := userReq.Validate(); err != nil {
		return usrResp, err
	}

	hashedPassword, err := util.HashPassword(userReq.Password)
	if err != nil {
		return usrResp, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to hash password:"), http.StatusInternalServerError, "internal error")
	}

	user, err := u.store.CreateUser(ctx, db.CreateUserParams{
		Username: userReq.Username,
		Password: hashedPassword,
		Email:    userReq.Email,
	})
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			return usrResp, util.NewErrorResponse(util.RequestError.New("user %q already exists", userReq.Username), http.StatusForbidden, fmt.Sprintf("user %q already exists", userReq.Username))
		}
		return usrResp, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create user"), http.StatusInternalServerError, "failed to create user")
	}

	usrResp = newCreateUserResponse(user)
	return usrResp, nil

}

func (u *UserServ) LoginUser(ctx context.Context, userReq LoginUserRequest) (LoginUserResponse, error) {

	if err := userReq.Validate(); err != nil {
		return LoginUserResponse{}, err
	}

	user, err := u.store.GetUser(ctx, userReq.Username)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return LoginUserResponse{}, util.NewErrorResponse(util.RequestError.New("%q was not found", userReq.Username), http.StatusNotFound, fmt.Sprintf("%q was not found", userReq.Username))
		}
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to login user"), http.StatusInternalServerError, "failed to login user")
	}

	if err = util.CheckPassword(userReq.Password, user.Password); err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(util.AuthenticationError.Wrap(err, "incorrect password"), http.StatusUnauthorized, "incorrect password")
	}

	accessToken, accessTokenPayload, err := u.tokenMaker.CreateToken(user.UserID, user.Role, u.config.AccessTokenDuration)
	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create access token"), http.StatusInternalServerError, "internal error")
	}

	refreshToken, refreshTokenPayload, err := u.tokenMaker.CreateToken(user.UserID, user.Role, u.config.RefreshTokenDuration)
	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create refresh token"), http.StatusInternalServerError, "internal error")
	}

	err = u.store.DeleteUserSession(ctx, user.UserID)
	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.New("failed to delete user session"), http.StatusInternalServerError, "internal error")
	}
	//TODO: hash or encrypt the refreshToken before storing it

	_, err = u.store.CreateUserSession(ctx, db.CreateUserSessionParams{
		SessionID:    refreshTokenPayload.ID,
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	})

	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create user session"), http.StatusInternalServerError, "internal error")
	}

	resp := LoginUserResponse{
		Username:              user.Username,
		Email:                 user.Email,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
	}

	return resp, nil

}
