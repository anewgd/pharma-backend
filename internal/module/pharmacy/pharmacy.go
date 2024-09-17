package pharmacy

import (
	"context"
	"regexp"

	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/module"
	"pharma-backend/internal/storage"
	"pharma-backend/platform/logger"
	"pharma-backend/platform/util"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type pharmacy struct {
	pharmacyPersistent storage.Pharmacy
	log                logger.Logger
}

func Init(log logger.Logger, pharmacyPersistent storage.Pharmacy) module.Pharmacy {
	return &pharmacy{
		pharmacyPersistent: pharmacyPersistent,
		log:                log,
	}
}

func (p *pharmacy) CreatePharmacy(ctx context.Context, pharmacyName string) (*dto.Pharmacy, error) {
	err := validation.Validate(pharmacyName, validation.Required.Error("pharmacy name is required"),
		validation.Length(3, 100).Error("must be between 3 and 100 characters long"),
		validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
		validation.Match(regexp.MustCompile("([A-Za-z])+")).Error("must contain atleast one or more letters"))

	if err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.log.Error(ctx, "validation failed", zap.Error(err), zap.String("input", pharmacyName))
		return nil, err
	}

	role, err := util.GetUserRole(ctx)
	if err != nil {
		return nil, err
	}

	if role != util.Admin {
		err = errors.ErrUnauthorized.New("insufficient permission. ADMIN role required")
		p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("user_role", role))
		return nil, err
	}

	userID, err := util.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	return p.pharmacyPersistent.CreatePharmacy(ctx, dto.CreatePharmacyRequest{
		PharmacyName: pharmacyName,
		UserID:       userID,
	})
}

func (p *pharmacy) CreatePharmacyBranch(ctx context.Context, param dto.CreatePharmacyBranchReq) (*dto.PharmacyBranch, error) {
	role, err := util.GetUserRole(ctx)
	if err != nil {
		return nil, err
	}

	if role != util.Admin {
		err = errors.ErrUnauthorized.New("insufficient permission. ADMIN role required")
		p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("user_role", role))
		return nil, err
	}

	userID, err := util.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	pharmacy, err := p.pharmacyPersistent.GetPharmacyByAdminID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return p.pharmacyPersistent.CreateBranch(ctx, dto.CreatePharmacyBranchRequest{
		PharmacyID:          pharmacy.PharmacyID,
		PharmacyBranchName:  param.PharmacyBranchName,
		City:                param.City,
		SubCity:             param.SubCity,
		SpecialLocationName: param.SpecialLocationName,
	})
}

func (p *pharmacy) CreatePharmacyBranchManager(ctx context.Context, param dto.CreatePharmacistReq) (*dto.Pharmacist, error) {
	role, err := util.GetUserRole(ctx)
	if err != nil {
		return nil, err
	}

	if role != util.Admin {
		err = errors.ErrUnauthorized.New("insufficient permission. ADMIN role required")
		p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("user_role", role))
		return nil, err
	}

	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	branch, err := p.pharmacyPersistent.GetPharmacyBranchByName(ctx, param.PharmacyBranchName)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "unable to hash password")
		p.log.Error(ctx, "unable to create user", zap.Error(err))
		return nil, err
	}

	return p.pharmacyPersistent.CreatePharmacist(ctx, dto.CreatePharmacistRequest{
		PharmacyBranchID: branch.PharmacyBranchID,
		Username:         param.Username,
		Password:         string(hashedPassword),
		Email:            param.Email,
		Role:             util.Manager,
	})

}
func (p *pharmacy) CreatePharmacist(ctx context.Context, param dto.CreatePharmacistReq) (*dto.Pharmacist, error) {
	role, err := util.GetUserRole(ctx)
	if err != nil {
		return nil, err
	}

	if role != util.Admin && role != util.Manager {
		err = errors.ErrUnauthorized.New("insufficient permission. ADMIN or MANAGER role required")
		p.log.Error(ctx, "unable to create pharmacy", zap.Error(err), zap.String("user_role", role))
		return nil, err
	}

	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	branch, err := p.pharmacyPersistent.GetPharmacyBranchByName(ctx, param.PharmacyBranchName)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "unable to hash password")
		p.log.Error(ctx, "unable to create user", zap.Error(err))
		return nil, err
	}

	return p.pharmacyPersistent.CreatePharmacist(ctx, dto.CreatePharmacistRequest{
		PharmacyBranchID: branch.PharmacyBranchID,
		Username:         param.Username,
		Password:         string(hashedPassword),
		Email:            param.Email,
		Role:             util.Pharmacist,
	})
}

func (p *pharmacy) GetPharmacist(ctx context.Context, param dto.LoginUserRequest) (*dto.Pharmacist, error) {
	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	pharmacist, err := p.pharmacyPersistent.GetPharmacistByUsername(ctx, param.Username)
	if err != nil {
		return nil, err
	}

	if err = util.CheckPassword(param.Password, pharmacist.Password); err != nil {
		err = errors.ErrUnauthorized.New("incorrect password")
		p.log.Error(ctx, "unable to get user", zap.Error(err))
		return nil, err
	}

	//TODO: Add tokens

	return &dto.Pharmacist{
		PharmacistID:     pharmacist.PharmacistID,
		PharmacyBranchID: pharmacist.PharmacyBranchID,
		Username:         pharmacist.Username,
		Email:            pharmacist.Email,
		Role:             pharmacist.Role,
	}, nil
}
