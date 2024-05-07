package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/util"
	"github.com/jackc/pgx/v5"
	"github.com/joomcode/errorx"
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
		return db.Drug{}, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get user role"), http.StatusInternalServerError, "internal error")
	}

	if userRole != util.Manager {
		return db.Drug{}, util.NewErrorResponse(util.AuthorizationError.New("insufficient user permissions"), http.StatusForbidden, "insufficient user permissions")
	}

	res := db.Drug{}
	if err := drugReq.Validate(); err != nil {
		return res, err
	}

	userID, err := util.GetUserID(ctx)
	if err != nil {
		return res, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to get user id"), http.StatusInternalServerError, "internal error")
	}

	fmt.Println(userID)

	manager, err := drugService.store.GetPharmacist(ctx, userID)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return res, util.NewErrorResponse(util.RequestError.New("manager not found"), http.StatusNotFound, "manager not found")
		}
		return res, util.NewErrorResponse(errorx.InternalError.New("failed to get manager info"), http.StatusInternalServerError, "internal error")
	}

	expDate, err := time.ParseInLocation("2006-02-01", drugReq.ExpirationDate, time.Local)
	if err != nil {
		return db.Drug{}, util.NewErrorResponse(util.RequestError.Wrap(err, "failed to parse expiration date"), http.StatusBadRequest, fmt.Sprintf("invalid date %q", drugReq.ExpirationDate))
	}
	manufacturingDate, err := time.ParseInLocation("2006-02-01", drugReq.ManufacturingDate, time.Local)
	if err != nil {
		return db.Drug{}, util.NewErrorResponse(util.RequestError.Wrap(err, "failed to parse manufacturing date"), http.StatusBadRequest, fmt.Sprintf("invalid date %q", drugReq.ManufacturingDate))
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
		if util.ErrorCode(err) == util.UniqueViolation {
			return res, util.NewErrorResponse(util.RequestError.New("drug already exists"), http.StatusForbidden, fmt.Sprintf("drug %q already exists", drug.BrandName))
		}
		return res, util.NewErrorResponse(errorx.InternalError.Wrap(err, "failed to create drug"), http.StatusInternalServerError, "internal error")
	}

	return drug, nil

}
