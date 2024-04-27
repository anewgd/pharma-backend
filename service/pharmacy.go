package service

import (
	"time"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/google/uuid"
)

type CreatePharmacyRequest struct {
	PharmacyName string `json:"pharmacy_name"`
}

type CreatePharmacyResponse struct {
	PharmacyID      uuid.UUID `json:"pharmacy_id"`
	PharmacyName    string    `json:"pharmacy_name"`
	PharmacyAdminID uuid.UUID `json:"pharmacy_admin_id"`
	CreatedAt       time.Time `json:"created_at"`
}

func newCreatePharmacyResponse(pharmacy db.Pharmacy) CreatePharmacyResponse {
	return CreatePharmacyResponse{
		PharmacyID:      pharmacy.PharmacyID,
		PharmacyName:    pharmacy.PharmacyName,
		PharmacyAdminID: pharmacy.UserID,
		CreatedAt:       pharmacy.CreatedAt,
	}
}

type CreatePharmacistRequest struct {
	PharmacyBranchID uuid.UUID `json:"pharmacy__id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
}

type CreatePharmacistResponse struct {
	PharmacyID uuid.UUID `json:"pharmacy_id"`
	UserID     string    `json:"user_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
}

type CreatePharmacyBranchRequest struct {
	PharmacyBranchName  string `json:"pharmacy_branch_name"`
	City                string `json:"city"`
	SubCity             string `json:"sub_city"`
	SpecialLocationName string `json:"special_location_name"`
}

type CreatePharmacyBranchResponse struct {
	PharmacyBranchID    uuid.UUID `json:"pharmacy_branch_id"`
	PharmacyID          uuid.UUID `json:"pharmacy_id"`
	PharmacyBranchName  string    `json:"pharmacy_branch_name"`
	City                string    `json:"city"`
	SubCity             string    `json:"sub_city"`
	SpecialLocationName string    `json:"special_location_name"`
}

type CreatePharmacyManagerRequest struct {
	PharmacyBranchName string `json:"pharmacy_branch_name"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Email              string `json:"email"`
}
