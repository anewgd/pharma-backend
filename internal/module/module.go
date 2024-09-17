package module

import (
	"context"

	"pharma-backend/internal/constants/model/dto"
)

type User interface {
	CreateUser(ctx context.Context, param dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetUser(ctx context.Context, param dto.LoginUserRequest) (*dto.LoginUserResponse, error)
}

type Pharmacy interface {
	CreatePharmacy(ctx context.Context, pharmacyName string) (*dto.Pharmacy, error)
	CreatePharmacyBranch(ctx context.Context, param dto.CreatePharmacyBranchReq) (*dto.PharmacyBranch, error)
	CreatePharmacyBranchManager(ctx context.Context, param dto.CreatePharmacistReq) (*dto.Pharmacist, error)
	CreatePharmacist(ctx context.Context, param dto.CreatePharmacistReq) (*dto.Pharmacist, error)
	GetPharmacist(ctx context.Context, param dto.LoginUserRequest) (*dto.Pharmacist, error)
}

type Drug interface {
	CreateDrug(ctx context.Context, param dto.CreateDrugReq) (*dto.Drug, error)
}
