package pharmacy

import (
	"context"

	"pharma-backend/internal/constants/dbinstance"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/db"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/storage"
	"pharma-backend/platform/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type pharmacy struct {
	db  dbinstance.DBInstance
	log logger.Logger
}

func Init(db dbinstance.DBInstance, log logger.Logger) storage.Pharmacy {

	return &pharmacy{
		db:  db,
		log: log,
	}
}

func (p *pharmacy) CreatePharmacy(ctx context.Context, param dto.CreatePharmacyRequest) (*dto.Pharmacy, error) {
	pharmacy, err := p.db.CreatePharmacy(ctx, db.CreatePharmacyParams{
		PharmacyName: param.PharmacyName,
		UserID:       param.UserID,
	})

	if err != nil {
		if errors.ErrorCode(err) == errors.UniqueViolation {
			err = errors.ErrDataExists.New("pharmacy already exists")
			p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("pharmacy_name", param.PharmacyName))
			return nil, err
		}
		err = errors.ErrWriteError.Wrap(err, "could not create pharmacy")
		p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("pharmacy_name", param.PharmacyName))
		return nil, err
	}

	return &dto.Pharmacy{
		PharmacyID:   pharmacy.PharmacyID,
		UserID:       pharmacy.UserID,
		PharmacyName: pharmacy.PharmacyName,
		CreatedAt:    pharmacy.CreatedAt,
	}, nil

}
func (p *pharmacy) CreatePharmacist(ctx context.Context, param dto.CreatePharmacistRequest) (*dto.Pharmacist, error) {
	pharmacist, err := p.db.CreatePharmacist(ctx, db.CreatePharmacistParams{
		PharmacyBranchID: param.PharmacyBranchID,
		Username:         param.Username,
		Password:         param.Password,
		Email:            param.Email,
		Role:             param.Role,
	})

	if err != nil {
		if errors.ErrorCode(err) == errors.UniqueViolation {
			err = errors.ErrDataExists.New("pharmacist already exists")
			p.log.Error(ctx, "unable to create pharmacist", zap.Error(err), zap.String("pharmacist_username", param.Username))
			return nil, err
		}
		err = errors.ErrWriteError.Wrap(err, "could not create pharmacist")
		p.log.Error(ctx, "unable to create pharmacy branch pharmacist", zap.Error(err), zap.String("pharmacist_username", param.Username))
		return nil, err
	}

	return &dto.Pharmacist{
		PharmacistID:     pharmacist.PharmacistID,
		PharmacyBranchID: pharmacist.PharmacyBranchID,
		Username:         pharmacist.Username,
		Email:            pharmacist.Email,
		Role:             pharmacist.Role,
		CreatedAt:        pharmacist.CreatedAt,
	}, nil
}
func (p *pharmacy) CreateBranch(ctx context.Context, param dto.CreatePharmacyBranchRequest) (*dto.PharmacyBranch, error) {
	branch, err := p.db.CreatePharmacyBranch(ctx, db.CreatePharmacyBranchParams{
		PharmacyID:          param.PharmacyID,
		PharmacyBranchName:  param.PharmacyBranchName,
		City:                param.City,
		SubCity:             param.SubCity,
		SpecialLocationName: param.SpecialLocationName,
	})

	if err != nil {
		if errors.ErrorCode(err) == errors.UniqueViolation {
			err = errors.ErrDataExists.New("pharmacy branch already exists")
			p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("pharmacy_branch_name", param.PharmacyBranchName))
			return nil, err
		}
		err = errors.ErrWriteError.Wrap(err, "could not create pharmacy branch")
		p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("pharmacy_name", param.PharmacyBranchName))
		return nil, err
	}

	return &dto.PharmacyBranch{
		PharmacyBranchID:    branch.PharmacyBranchID,
		PharmacyID:          branch.PharmacyID,
		PharmacyBranchName:  branch.PharmacyBranchName,
		City:                branch.City,
		SubCity:             branch.SubCity,
		SpecialLocationName: branch.SpecialLocationName,
		CreatedAt:           branch.CreatedAt,
	}, nil
}

// func (p *pharmacy) CreateBranchManager(ctx context.Context, param dto.CreatePharmacyManagerRequest) (*dto.Pharmacist, error) {
// 	manager, err := p.db.CreatePharmacist(ctx, db.CreatePharmacistParams{
// 		PharmacyBranchID: param.PharmacBranchID,
// 		Username:         param.Username,
// 		Password:         param.Password,
// 		Email:            param.Email,
// 		Role:             "MANAGER", //Magic value, make it better
// 	})

// 	if err != nil {
// 		if errors.ErrorCode(err) == errors.UniqueViolation {
// 			err = errors.ErrDataExists.New("manager already exists")
// 			p.log.Error(ctx, "unable to create pharmacy manager", zap.Error(err), zap.String("pharmacy_manager_name", param.Username))
// 			return nil, err
// 		}
// 		err = errors.ErrWriteError.Wrap(err, "could not create pharmacy branch manager")
// 		p.log.Error(ctx, "unable to create pharmacy branch manager", zap.Error(err), zap.String("pharmacy_manager_name", param.Username))
// 		return nil, err
// 	}

// 	return &dto.Pharmacist{
// 		PharmacistID: manager.PharmacistID,
// 		PharmacyBranchID: manager.PharmacyBranchID,
// 		Username: manager.Username,
// 		Email: manager.Email,
// 		Role: manager.Role,
// 		CreatedAt: manager.CreatedAt,
// 	}, nil
// }

func (p *pharmacy) GetPharmacyByAdminID(ctx context.Context, userID uuid.UUID) (*dto.Pharmacy, error) {
	pharmacy, err := p.db.GetPharmacyByAdminID(ctx, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrNoRecordFound.New("pharmacy not found")
			p.log.Error(ctx, "unable to retrieve pharamcy", zap.Error(err), zap.String("user_id", userID.String()))
			return nil, err
		}
		err = errors.ErrReadError.Wrap(err, "unable to retrieve pharmacy")
		p.log.Error(ctx, "unable to retrieve pharamcy", zap.Error(err), zap.String("user_id", userID.String()))
		return nil, err
	}
	return &dto.Pharmacy{
		PharmacyID:   pharmacy.PharmacyID,
		UserID:       pharmacy.UserID,
		PharmacyName: pharmacy.PharmacyName,
		CreatedAt:    pharmacy.CreatedAt,
	}, nil

}

func (p *pharmacy) GetPharmacyBranchByName(ctx context.Context, branchName string) (*dto.PharmacyBranch, error) {
	branch, err := p.db.GetPharmacyBrachByName(ctx, branchName)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrNoRecordFound.New("pharmacy branch not found")
			p.log.Error(ctx, "unable to retrieve pharamcy branch", zap.Error(err), zap.String("branch_name", branchName))
			return nil, err
		}
		err = errors.ErrReadError.Wrap(err, "unable to retrieve pharmacy branch")
		p.log.Error(ctx, "unable to retrieve pharamcy branch", zap.Error(err), zap.String("branch_name", branchName))
		return nil, err
	}

	return &dto.PharmacyBranch{
		PharmacyBranchID:    branch.PharmacyBranchID,
		PharmacyID:          branch.PharmacyID,
		PharmacyBranchName:  branch.PharmacyBranchName,
		City:                branch.City,
		SubCity:             branch.SubCity,
		SpecialLocationName: branch.SpecialLocationName,
		CreatedAt:           branch.CreatedAt,
	}, nil
}

func (p *pharmacy) GetPharmacistByUsername(ctx context.Context, username string) (*dto.Pharmacist, error) {
	pharmacist, err := p.db.GetPharmacistByUsername(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrNoRecordFound.New("pharmacist not found")
			p.log.Error(ctx, "unable to retrieve pharamcist", zap.Error(err), zap.String("username", username))
			return nil, err
		}
		err = errors.ErrReadError.Wrap(err, "unable to retrieve pharmacy branch")
		p.log.Error(ctx, "unable to retrieve pharamcist", zap.Error(err), zap.String("username", username))
		return nil, err
	}

	return &dto.Pharmacist{
		PharmacistID:     pharmacist.PharmacistID,
		PharmacyBranchID: pharmacist.PharmacyBranchID,
		Username:         pharmacist.Username,
		Password:         pharmacist.Password,
		Email:            pharmacist.Email,
		Role:             pharmacist.Role,
		CreatedAt:        pharmacist.CreatedAt,
	}, nil
}

func (p *pharmacy) GetPharmacist(ctx context.Context, userID uuid.UUID) (*dto.Pharmacist, error) {
	pharmacist, err := p.db.GetPharmacist(ctx, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrNoRecordFound.New("pharmacist not found")
			p.log.Error(ctx, "unable to retrieve pharamcist", zap.Error(err), zap.String("user_id", userID.String()))
			return nil, err
		}
		err = errors.ErrReadError.Wrap(err, "unable to retrieve pharmacy branch")
		p.log.Error(ctx, "unable to retrieve pharamcist", zap.Error(err), zap.String("user_id", userID.String()))
		return nil, err
	}

	return &dto.Pharmacist{
		PharmacistID:     pharmacist.PharmacistID,
		PharmacyBranchID: pharmacist.PharmacyBranchID,
		Username:         pharmacist.Username,
		Email:            pharmacist.Email,
		Role:             pharmacist.Role,
		CreatedAt:        pharmacist.CreatedAt,
	}, nil
}
