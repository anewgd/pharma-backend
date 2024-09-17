package dto

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Drug struct {
	DrugID            uuid.UUID `json:"drug_id"`
	PharmacyBranchID  uuid.UUID `json:"pharmacy_branch_id"`
	BrandName         string    `json:"brand_name"`
	GenericName       string    `json:"generic_name"`
	Quantity          int64     `json:"quantity"`
	ExpirationDate    time.Time `json:"expiration_date"`
	ManufacturingDate time.Time `json:"manufacturing_date"`
	PharmacistID      uuid.UUID `json:"pharmacist_id"`
	AddedAt           time.Time `json:"added_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	DeletedAt         time.Time `json:"deleted_at,omitempty"`
}

type CreateDrugReq struct {
	BrandName         string `json:"brand_name"`
	GenericName       string `json:"generic_name"`
	Quantity          int64  `json:"quantity"`
	ExpirationDate    string `json:"expiration_date"`
	ManufacturingDate string `json:"manufacturing_date"`
}

type CreateDrugRequest struct {
	PharmacyBranchID  uuid.UUID `json:"pharmacy_branch_id"`
	BrandName         string    `json:"brand_name"`
	GenericName       string    `json:"generic_name"`
	Quantity          int64     `json:"quantity"`
	ExpirationDate    time.Time `json:"expiration_date"`
	ManufacturingDate time.Time `json:"manufacturing_date"`
	PharmacistID      uuid.UUID `json:"pharmacist_id"`
}

func (c CreateDrugReq) Validate() error {
	// expDate := c.ExpirationDate.String()

	return validation.ValidateStruct(&c,
		validation.Field(&c.BrandName,
			validation.Required,
			validation.Length(4, 30),
			validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
		),
		validation.Field(&c.GenericName,
			validation.Required,
			validation.Length(4, 30),
			validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
		),
		validation.Field(&c.Quantity,
			validation.Required,
			validation.Max(1000000),
			validation.Min(1),
		),
		validation.Field(&c.ExpirationDate,
			validation.Required,
			validation.Date("2006-01-02"),
		),

		validation.Field(&c.ManufacturingDate,
			validation.Date("2006-01-02"),
		),
	)
}
