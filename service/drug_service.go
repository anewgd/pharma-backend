package service

import (
	"context"

	db "github.com/anewgd/pharma_backend/data/sqlc"
)

type DrugServ struct {
	store db.Store
}

func NewDrugService(store db.Store) *DrugServ {
	return &DrugServ{
		store: store,
	}
}

func (drugService *DrugServ) AddDrug(ctx context.Context, drugData CreateDrugRequest) (*CreateDrugResponse, error) {

	resp := CreateDrugResponse{}
	// arg := db.CreateDrugParams{
	// 	PharmacyID:        ,
	// 	BrandName:         drugData.BrandName,
	// 	GenericName:       drugData.GenericName,
	// 	Quantity:          drugData.Quantity,
	// 	ExpirationDate:    drugData.ExpirationDate,
	// 	ManufacturingDate: drugData.ManufacturingDate,
	// 	UserID:            drugData.UserID,
	// }
	// drug, err := drugService.store.CreateDrug(ctx, arg)
	// if err != nil {
	// 	return &resp, err
	// }

	// resp = CreateDrugResponse{
	// 	DrugID:            drug.DrugID,
	// 	PharmacyID:        drug.PharmacyID,
	// 	GenericName:       drug.GenericName,
	// 	BrandName:         drug.BrandName,
	// 	Quantity:          drug.Quantity,
	// 	ExpirationDate:    drug.ExpirationDate,
	// 	ManufacturingDate: drug.ManufacturingDate,
	// 	UserID:            drug.UserID,
	// 	AddedAt:           drug.AddedAt,
	// }
	return &resp, nil
}
