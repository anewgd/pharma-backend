package drug

import (
	"context"
	"time"

	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/module"
	"pharma-backend/internal/storage"
	"pharma-backend/platform/logger"
	"pharma-backend/platform/util"

	"go.uber.org/zap"
)

type drug struct {
	drugPersistent     storage.Drug
	pharmacyPersistent storage.Pharmacy
	log                logger.Logger
}

func Init(log logger.Logger, drugPersistent storage.Drug, pharmacyPercistent storage.Pharmacy) module.Drug {
	return &drug{
		drugPersistent:     drugPersistent,
		pharmacyPersistent: pharmacyPercistent,
		log:                log,
	}
}

func (d *drug) CreateDrug(ctx context.Context, param dto.CreateDrugReq) (*dto.Drug, error) {
	role, err := util.GetUserRole(ctx)
	if err != nil {
		return nil, err
	}

	if role != util.Manager {
		err = errors.ErrUnauthorized.New("insufficient permission. MANAGER role required")
		d.log.Error(ctx, "unable to create drug", zap.Error(err), zap.String("user_role", role))
		return nil, err
	}

	userID, err := util.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		d.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	manager, err := d.pharmacyPersistent.GetPharmacist(ctx, userID)
	if err != nil {
		return nil, err
	}

	expirationDate, err := time.ParseInLocation("2006-02-01", param.ExpirationDate, time.Local)
	if err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid expiration date")
		d.log.Error(ctx, "failed to parse expiration date", zap.Error(err))
		return nil, err
	}

	manufacturingDate, err := time.ParseInLocation("2006-02-01", param.ExpirationDate, time.Local)
	if err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid manufacturing date")
		d.log.Error(ctx, "failed to parse manufacturing date", zap.Error(err))
		return nil, err
	}

	return d.drugPersistent.Create(ctx, dto.CreateDrugRequest{
		PharmacyBranchID:  manager.PharmacyBranchID,
		BrandName:         param.BrandName,
		GenericName:       param.GenericName,
		Quantity:          int64(param.Quantity),
		ExpirationDate:    expirationDate,
		ManufacturingDate: manufacturingDate,
		PharmacistID:      manager.PharmacistID,
	})

}
