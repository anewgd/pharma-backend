package service

import (
	"context"
	"fmt"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/util"
	"github.com/anewgd/pharma_backend/util/token"
	"github.com/jackc/pgx/v5"
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
		return nil, err
	}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
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
		return usrResp, err
	}

	user, err := u.store.CreateUser(ctx, db.CreateUserParams{
		Username: userReq.Username,
		Password: hashedPassword,
		Email:    userReq.Email,
	})
	if err != nil {
		return usrResp, err
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
			return LoginUserResponse{}, fmt.Errorf("user not found")
		}
		return LoginUserResponse{}, err
	}

	if err = util.CheckPassword(userReq.Password, user.Password); err != nil {
		return LoginUserResponse{}, fmt.Errorf("incorrect password")
	}

	accessToken, accessTokenPayload, err := u.tokenMaker.CreateToken(user.UserID, user.Role, u.config.AccessTokenDuration)
	if err != nil {
		return LoginUserResponse{}, err
	}

	refreshToken, refreshTokenPayload, err := u.tokenMaker.CreateToken(user.UserID, user.Role, u.config.RefreshTokenDuration)
	if err != nil {
		return LoginUserResponse{}, err
	}

	err = u.store.DeleteUserSession(ctx, user.UserID)
	if err != nil {
		return LoginUserResponse{}, err
	}
	//TODO: hash or encrypt the refreshToken before storing it

	_, err = u.store.CreateUserSession(ctx, db.CreateUserSessionParams{
		SessionID:    refreshTokenPayload.ID,
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	})

	if err != nil {
		return LoginUserResponse{}, err
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
