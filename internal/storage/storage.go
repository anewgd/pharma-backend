package storage

import (
	"context"

	"pharma-backend/internal/constants/model/dto"

	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, param dto.CreateUserRequest) (*dto.User, error)
	Get(ctx context.Context, username string) (*dto.User, error)
	CheckUserExists(ctx context.Context, email string) (bool, error)
}

type UserSession interface {
	Create(ctx context.Context, param dto.CreateUserSession) (*dto.UserSession, error)
	Get(ctx context.Context, sessionID uuid.UUID) (*dto.UserSession, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type PharmacySession interface {
	Create(ctx context.Context, param dto.CreatePharmacistSession) (*dto.PharmacistSession, error)
	Get(ctx context.Context, sessionID uuid.UUID) (*dto.PharmacistSession, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type Drug interface {
	Create(ctx context.Context, param dto.CreateDrugRequest) (*dto.Drug, error)
}

type Pharmacy interface {
	CreatePharmacy(ctx context.Context, param dto.CreatePharmacyRequest) (*dto.Pharmacy, error)
	CreatePharmacist(ctx context.Context, param dto.CreatePharmacistRequest) (*dto.Pharmacist, error)
	CreateBranch(ctx context.Context, param dto.CreatePharmacyBranchRequest) (*dto.PharmacyBranch, error)
	// CreateBranchManager(ctx context.Context, param dto.CreatePharmacyManagerRequest) (*dto.Pharmacist, error)
	GetPharmacyByAdminID(ctx context.Context, userID uuid.UUID) (*dto.Pharmacy, error)
	GetPharmacyBranchByName(ctx context.Context, branchName string) (*dto.PharmacyBranch, error)
	GetPharmacistByUsername(ctx context.Context, username string) (*dto.Pharmacist, error)
	GetPharmacist(ctx context.Context, userID uuid.UUID) (*dto.Pharmacist, error)
}
