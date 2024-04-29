package service

import (
	"context"
	"fmt"
	"time"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/util"
)

type DrugServ struct {
	store db.Store
}

func NewDrugService(store db.Store) *DrugServ {
	return &DrugServ{
		store: store,
	}
}

func (drugService *DrugServ) AddDrug(ctx context.Context, drugReq CreateDrugRequest) (db.Drug, error) {

	userRole, err := util.GetUserRole(ctx)

	if err != nil {
		return db.Drug{}, err
	}

	if userRole != util.Manager {
		return db.Drug{}, fmt.Errorf("insufficient permissions")
	}

	res := db.Drug{}
	if err := drugReq.Validate(); err != nil {
		return res, err
	}

	role, err := util.GetUserRole(ctx)
	if err != nil {
		return res, err
	}
	if role != util.Manager {
		return res, fmt.Errorf("unauthorized role")
	}

	userID, err := util.GetUserID(ctx)
	if err != nil {
		return res, err
	}

	fmt.Println(userID)

	manager, err := drugService.store.GetPharmacist(ctx, userID)
	if err != nil {
		return res, err
	}

	expDate, err := time.ParseInLocation("2006-02-01", drugReq.ExpirationDate, time.Local)
	if err != nil {
		return db.Drug{}, err
	}
	manufacturingDate, err := time.ParseInLocation("2006-02-01", drugReq.ManufacturingDate, time.Local)
	if err != nil {
		return db.Drug{}, err
	}

	drug, err := drugService.store.CreateDrug(ctx, db.CreateDrugParams{
		PharmacyBranchID:  manager.PharmacyBranchID,
		BrandName:         drugReq.BrandName,
		GenericName:       drugReq.GenericName,
		Quantity:          int64(drugReq.Quantity),
		ExpirationDate:    expDate,
		ManufacturingDate: manufacturingDate,
		PharmacistID:      manager.PharmacistID,
	})
	if err != nil {
		return res, err
	}

	return drug, nil

}
