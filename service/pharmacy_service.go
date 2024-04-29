package service

import (
	"context"
	"fmt"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/util"
	"github.com/anewgd/pharma_backend/util/token"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PharmacyService interface {
	CreatePharmacy(ctx context.Context, pharmaReq CreatePharmacyRequest) (CreatePharmacyResponse, error)
	CreatePharmacyBranch(ctx context.Context, pharmaReq CreatePharmacyBranchRequest) (db.PharmacyBranch, error)
	CreateBranchManager(ctx context.Context, pharmaReq CreatePharmacyManagerRequest) (db.Pharmacist, error)
	PharmacyLogin(ctx context.Context, pharmaReq LoginUserRequest) (LoginUserResponse, error)
}

type PharmacyServ struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewPharmacyService(store db.Store) (*PharmacyServ, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		return nil, err
	}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	return &PharmacyServ{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}, nil
}

func (p *PharmacyServ) CreatePharmacy(ctx context.Context, pharmaReq CreatePharmacyRequest) (CreatePharmacyResponse, error) {

	pharmaResp := CreatePharmacyResponse{}
	userID := ctx.Value(util.UserID)

	id, ok := (userID).(uuid.UUID)
	if !ok {
		return pharmaResp, fmt.Errorf("cannot find user id")
	}
	if err := pharmaReq.Validate(); err != nil {
		return pharmaResp, err
	}

	pharmacy, err := p.store.CreatePharmacy(ctx, db.CreatePharmacyParams{
		PharmacyName: pharmaReq.PharmacyName,
		UserID:       id,
	})
	if err != nil {
		return pharmaResp, err
	}

	pharmaResp = newCreatePharmacyResponse(pharmacy)
	return pharmaResp, nil
}

func (p *PharmacyServ) CreatePharmacyBranch(ctx context.Context, pharmaReq CreatePharmacyBranchRequest) (db.PharmacyBranch, error) {
	res := db.PharmacyBranch{}
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return res, err
	}

	if err := pharmaReq.Validate(); err != nil {
		return res, err
	}

	pharmacy, err := p.store.GetPharmacyByAdminID(ctx, userID)
	if err != nil {
		return res, err
	}

	branch, err := p.store.CreatePharmacyBranch(ctx, db.CreatePharmacyBranchParams{
		PharmacyID:          pharmacy.PharmacyID,
		PharmacyBranchName:  pharmaReq.PharmacyBranchName,
		City:                pharmaReq.City,
		SubCity:             pharmaReq.SubCity,
		SpecialLocationName: pharmaReq.SpecialLocationName,
	})
	if err != nil {
		return res, err
	}
	res = branch

	return res, nil

}

func (p *PharmacyServ) CreateBranchManager(ctx context.Context, pharmaReq CreatePharmacyManagerRequest) (db.Pharmacist, error) {
	res := db.Pharmacist{}
	// // userID, err := util.GetUserID(ctx)
	// if err != nil {
	// 	return res, err
	// }

	if err := pharmaReq.Validate(); err != nil {
		return res, err
	}

	branch, err := p.store.GetPharmacyBrachByName(ctx, pharmaReq.PharmacyBranchName)
	if err != nil {
		return res, err
	}

	hashedPassword, err := util.HashPassword(pharmaReq.Password)
	if err != nil {
		return res, fmt.Errorf("failed to hash password")
	}

	manager, err := p.store.CreatePharmacist(ctx, db.CreatePharmacistParams{
		PharmacyBranchID: branch.PharmacyBranchID,
		Username:         pharmaReq.Username,
		Password:         hashedPassword,
		Email:            pharmaReq.Email,
		Role:             util.Manager,
	})
	if err != nil {
		return res, err
	}
	return manager, nil
}

func (p *PharmacyServ) PharmacyLogin(ctx context.Context, pharmaReq LoginUserRequest) (LoginUserResponse, error) {
	if err := pharmaReq.Validate(); err != nil {
		return LoginUserResponse{}, err
	}

	pharmacist, err := p.store.GetPharmacistByUsername(ctx, pharmaReq.Username)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return LoginUserResponse{}, fmt.Errorf("user not found")
		}
		return LoginUserResponse{}, err
	}

	if err = util.CheckPassword(pharmaReq.Password, pharmacist.Password); err != nil {
		return LoginUserResponse{}, fmt.Errorf("incorrect password")
	}

	accessToken, accessTokenPayload, err := p.tokenMaker.CreateToken(pharmacist.PharmacistID, pharmacist.Role, p.config.AccessTokenDuration)
	if err != nil {
		return LoginUserResponse{}, err
	}

	refreshToken, refreshTokenPayload, err := p.tokenMaker.CreateToken(pharmacist.PharmacistID, pharmacist.Role, p.config.RefreshTokenDuration)
	if err != nil {
		return LoginUserResponse{}, err
	}

	err = p.store.DeletePharmacistSession(ctx, pharmacist.PharmacistID)
	if err != nil {
		return LoginUserResponse{}, err
	}
	//TODO: hash or encrypt the refreshToken before storing it

	_, err = p.store.CreatePharmacistSession(ctx, db.CreatePharmacistSessionParams{
		SessionID:    refreshTokenPayload.ID,
		PharmacistID: pharmacist.PharmacistID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	})

	if err != nil {
		return LoginUserResponse{}, err
	}

	resp := LoginUserResponse{
		Username:              pharmacist.Username,
		Email:                 pharmacist.Email,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
	}

	return resp, nil
}
