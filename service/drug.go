package service

import (
	"context"

	db "github.com/anewgd/pharma_backend/data/sqlc"
)

type DrugService interface {
	AddDrug(ctx context.Context, drugData CreateDrugRequest) (db.Drug, error)
}

type CreateDrugRequest struct {
	BrandName         string `json:"brand_name"`
	GenericName       string `json:"generic_name"`
	Quantity          int    `json:"quantity"`
	ExpirationDate    string `json:"expiration_date" binding:"required"`
	ManufacturingDate string `json:"manufacturing_date" binding:"required"`
}
