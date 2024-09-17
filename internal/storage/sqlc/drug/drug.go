package drug

import (
	"context"

	"pharma-backend/internal/constants/dbinstance"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/db"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/storage"
	"pharma-backend/platform/logger"

	"go.uber.org/zap"
)

type drug struct {
	db  dbinstance.DBInstance
	log logger.Logger
}

func Init(db dbinstance.DBInstance, log logger.Logger) storage.Drug {

	return &drug{
		db:  db,
		log: log,
	}
}

func (d *drug) Create(ctx context.Context, param dto.CreateDrugRequest) (*dto.Drug, error) {
	drug, err := d.db.CreateDrug(ctx, db.CreateDrugParams{
		PharmacyBranchID:  param.PharmacyBranchID,
		BrandName:         param.BrandName,
		GenericName:       param.GenericName,
		Quantity:          int64(param.Quantity),
		ExpirationDate:    param.ExpirationDate,
		ManufacturingDate: param.ManufacturingDate,
		PharmacistID:      param.PharmacistID,
	})

	if err != nil {
		if errors.ErrorCode(err) == errors.UniqueViolation {
			err = errors.ErrDataExists.New("drug already exists")
			d.log.Error(ctx, "unable to create drug", zap.Error(err), zap.Any("input", param))
			return nil, err
		}
		err = errors.ErrWriteError.Wrap(err, "could not create drug")
		d.log.Error(ctx, "unable to create drug", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	return &dto.Drug{
		DrugID:            drug.DrugID,
		PharmacyBranchID:  drug.PharmacyBranchID,
		BrandName:         drug.BrandName,
		GenericName:       drug.GenericName,
		Quantity:          drug.Quantity,
		ExpirationDate:    drug.ExpirationDate,
		ManufacturingDate: drug.ManufacturingDate,
		PharmacistID:      drug.PharmacistID,
		AddedAt:           drug.AddedAt,
	}, nil
}
