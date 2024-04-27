package service

import (
	"context"
	"fmt"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/util"
	"github.com/google/uuid"
)

type PharmacyService interface {
	CreatePharmacy(ctx context.Context, pharmaReq CreatePharmacyRequest) (CreatePharmacyResponse, error)
	CreatePharmacyBranch(ctx context.Context, pharmaReq CreatePharmacyBranchRequest) (db.PharmacyBranch, error)
	CreateBranchManager(ctx context.Context, pharmaReq CreatePharmacyManagerRequest) (db.Pharmacist, error)
}

type PharmacyServ struct {
	store db.Store
}

func NewPharmacyService(store db.Store) *PharmacyServ {
	return &PharmacyServ{
		store: store,
	}
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
