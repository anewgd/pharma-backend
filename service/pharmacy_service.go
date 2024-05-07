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
		return nil, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to load configuration file"), http.StatusInternalServerError, "internal server error")
	}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create token maker"), http.StatusInternalServerError, "internal error")
	}
	return &PharmacyServ{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}, nil
}

func (p *PharmacyServ) CreatePharmacy(ctx context.Context, pharmaReq CreatePharmacyRequest) (CreatePharmacyResponse, error) {

	role, err := util.GetUserRole(ctx)
	if err != nil {
		return CreatePharmacyResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get user role"), http.StatusInternalServerError, "interal error")
	}
	if role != util.Admin {
		return CreatePharmacyResponse{}, util.NewErrorResponse(util.AuthorizationError.New("insufficient user permissions"), http.StatusForbidden, "insufficient user permissions")
	}

	pharmaResp := CreatePharmacyResponse{}
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return pharmaResp, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get user id"), http.StatusInternalServerError, "interal error")
	}
	if err := pharmaReq.Validate(); err != nil {
		return pharmaResp, err
	}

	pharmacy, err := p.store.CreatePharmacy(ctx, db.CreatePharmacyParams{
		PharmacyName: pharmaReq.PharmacyName,
		UserID:       userID,
	})
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			return pharmaResp, util.NewErrorResponse(util.RequestError.New("pharmacy %q already exists", pharmaReq.PharmacyName), http.StatusForbidden, fmt.Sprintf("pharmacy %q already exists", pharmaReq.PharmacyName))
		}
		return pharmaResp, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create pharmacy"), http.StatusInternalServerError, "internal error")
	}

	pharmaResp = newCreatePharmacyResponse(pharmacy)
	return pharmaResp, nil
}

func (p *PharmacyServ) CreatePharmacyBranch(ctx context.Context, pharmaReq CreatePharmacyBranchRequest) (db.PharmacyBranch, error) {

	role, err := util.GetUserRole(ctx)
	if err != nil {
		return db.PharmacyBranch{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get user role"), http.StatusInternalServerError, "interal error")
	}
	if role != util.Admin {
		return db.PharmacyBranch{}, util.NewErrorResponse(util.AuthorizationError.New("insufficient user permissions"), http.StatusForbidden, "insufficient user permissions")
	}

	res := db.PharmacyBranch{}
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return res, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get user id"), http.StatusInternalServerError, "interal error")
	}

	if err := pharmaReq.Validate(); err != nil {
		return res, err
	}

	pharmacy, err := p.store.GetPharmacyByAdminID(ctx, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return res, util.NewErrorResponse(util.RequestError.New("pharmacy was not found"), http.StatusNotFound, "no pharmacy found associated with user")
		}
		return res, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get pharmacy"), http.StatusInternalServerError, "internal error")
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
	role, err := util.GetUserRole(ctx)
	if err != nil {
		return res, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get user role"), http.StatusInternalServerError, "interal error")
	}
	if role != util.Admin {
		return res, util.NewErrorResponse(util.AuthorizationError.New("insufficient user permissions"), http.StatusForbidden, "insufficient user permissions")
	}

	if err := pharmaReq.Validate(); err != nil {
		return res, err
	}

	branch, err := p.store.GetPharmacyBrachByName(ctx, pharmaReq.PharmacyBranchName)
	if err != nil {
		if err == pgx.ErrNoRows {
			return res, util.NewErrorResponse(
				util.RequestError.New("pharmacy branch %q was not found",
					pharmaReq.PharmacyBranchName), http.StatusNotFound,
				fmt.Sprintf("pharmacy branch %q was not found", pharmaReq.PharmacyBranchName))
		}
		return res, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get pharmacy"), http.StatusInternalServerError, "internal error")
	}

	hashedPassword, err := util.HashPassword(pharmaReq.Password)
	if err != nil {
		return res, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to hash password:"), http.StatusInternalServerError, "internal error")
	}

	manager, err := p.store.CreatePharmacist(ctx, db.CreatePharmacistParams{
		PharmacyBranchID: branch.PharmacyBranchID,
		Username:         pharmaReq.Username,
		Password:         hashedPassword,
		Email:            pharmaReq.Email,
		Role:             util.Manager,
	})
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			return db.Pharmacist{}, util.NewErrorResponse(util.RequestError.New("manager %q already exists", pharmaReq.Username), http.StatusForbidden, fmt.Sprintf("manager %q already exists", pharmaReq.Username))
		}
		return db.Pharmacist{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create manager"), http.StatusInternalServerError, "failed to create manager")
	}
	return manager, nil
}

func (p *PharmacyServ) PharmacyLogin(ctx context.Context, pharmaReq LoginUserRequest) (LoginUserResponse, error) {

	//TODO: Check the role of the user
	if err := pharmaReq.Validate(); err != nil {
		return LoginUserResponse{}, err
	}

	pharmacist, err := p.store.GetPharmacistByUsername(ctx, pharmaReq.Username)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return LoginUserResponse{}, util.NewErrorResponse(util.RequestError.New("%q was not found", pharmaReq.Username), http.StatusNotFound, fmt.Sprintf("%q was not found", pharmaReq.Username))
		}
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to login user"), http.StatusInternalServerError, "internal error")
	}

	if err = util.CheckPassword(pharmaReq.Password, pharmacist.Password); err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(util.AuthenticationError.Wrap(err, "incorrect password"), http.StatusUnauthorized, "incorrect password")
	}

	accessToken, accessTokenPayload, err := p.tokenMaker.CreateToken(pharmacist.PharmacistID, pharmacist.Role, p.config.AccessTokenDuration)
	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create access token"), http.StatusInternalServerError, "internal error")
	}

	refreshToken, refreshTokenPayload, err := p.tokenMaker.CreateToken(pharmacist.PharmacistID, pharmacist.Role, p.config.RefreshTokenDuration)
	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create refresh token"), http.StatusInternalServerError, "internal error")
	}

	err = p.store.DeletePharmacistSession(ctx, pharmacist.PharmacistID)
	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.New("failed to delete pharmacist session"), http.StatusInternalServerError, "internal error")
	}
	//TODO: hash or encrypt the refreshToken before storing it

	_, err = p.store.CreatePharmacistSession(ctx, db.CreatePharmacistSessionParams{
		SessionID:    refreshTokenPayload.ID,
		PharmacistID: pharmacist.PharmacistID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	})

	if err != nil {
		return LoginUserResponse{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create pharmacist session"), http.StatusInternalServerError, "internal error")
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
