package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type DrugService interface {
	AddDrug(ctx context.Context, drugData CreateDrugRequest) (*CreateDrugResponse, error)
}

type CreateDrugRequest struct {
	BrandName         string    `json:"brand_name"`
	GenericName       string    `json:"generic_name"`
	Quantity          int64     `json:"quantity"`
	ExpirationDate    time.Time `json:"expiration_date"`
	ManufacturingDate time.Time `json:"manufacturing_date"`
}
type CreateDrugResponse struct {
	DrugID            uuid.UUID `json:"drug_id"`
	PharmacyID        uuid.UUID `json:"pharmacy_id"`
	BrandName         string    `json:"brand_name"`
	GenericName       string    `json:"generic_name"`
	Quantity          int64     `json:"quantity"`
	ExpirationDate    time.Time `json:"expiration_date"`
	ManufacturingDate time.Time `json:"manufacturing_date"`
	// UserID            uuid.UUID `json:"user_id"`
	AddedAt           time.Time `json:"added_at"`
}
