package dto

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Pharmacy struct {
	PharmacyID   uuid.UUID `json:"pharmacy_id,omitempty"`
	PharmacyName string    `json:"pharmacy_name"`
	UserID       uuid.UUID `json:"user_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

type CreatePharmacyReq struct {
	PharmacyName string `json:"pharmacy_name"`
}

type CreatePharmacyRequest struct {
	PharmacyName string    `json:"pharmacy_name"`
	UserID       uuid.UUID `json:"user_id"`
}

func (pharmaReq CreatePharmacyRequest) Validate() error {
	return validation.ValidateStruct(&pharmaReq,
		validation.Field(&pharmaReq.PharmacyName,
			validation.Required.Error("pharmacy name is required"),
			validation.Length(3, 100).Error("must be between 3 and 100 characters long"),
			validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
			validation.Match(regexp.MustCompile("([A-Za-z])+")).Error("must contain atleast one or more letters"),
		),
	)
}

type PharmacyBranch struct {
	PharmacyBranchID    uuid.UUID `json:"pharmacy_branch_id"`
	PharmacyID          uuid.UUID `json:"pharmacy_id"`
	PharmacyBranchName  string    `json:"pharmacy_branch_name"`
	City                string    `json:"city"`
	SubCity             string    `json:"sub_city"`
	SpecialLocationName string    `json:"special_location_name"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	DeletedAt           time.Time `json:"deleted_at,omitempty"`
}

type CreatePharmacyBranchReq struct {
	PharmacyBranchName  string `json:"pharmacy_branch_name"`
	City                string `json:"city"`
	SubCity             string `json:"sub_city"`
	SpecialLocationName string `json:"special_location_name"`
}
type CreatePharmacyBranchRequest struct {
	PharmacyID          uuid.UUID `json:"pharmacy_id"`
	PharmacyBranchName  string    `json:"pharmacy_branch_name"`
	City                string    `json:"city"`
	SubCity             string    `json:"sub_city"`
	SpecialLocationName string    `json:"special_location_name"`
}

func (pharmaReq CreatePharmacyBranchReq) Validate() error {
	return validation.ValidateStruct(&pharmaReq,
		validation.Field(&pharmaReq.PharmacyBranchName,
			validation.Required,
			validation.Length(5, 100),
			validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
			validation.Match(regexp.MustCompile("([A-Za-z])+")).Error("must contain atleast one or more letters"),
		),
		validation.Field(&pharmaReq.City,
			validation.Required,
			validation.Length(2, 100),
			validation.Match(regexp.MustCompile(`([A-Za-z\s])+`)).Error("must contain only letters"),
			// validation.Match(regexp.MustCompile("^[A-Z]+")).Error("name must start with an uppercase letter"),
			// validation.Match(regexp.MustCompile("([A-Za-z])+")).Error("name must end with a letter"),
			// validation.Match(regexp.MustCompile(`((([a-zA-Z0-9])+([\s]?))+)`)).Error("invalid city name"),

		),
		validation.Field(&pharmaReq.SubCity,
			validation.Required,
			validation.Length(2, 100),
			validation.Match(regexp.MustCompile(`([A-Za-z\s])+`)).Error("must contain only letters"),
			// validation.Match(regexp.MustCompile("^[A-Z]+")).Error("name must start with an uppercase letter"),
			// validation.Match(regexp.MustCompile("[A-Za-z]+$")).Error("name must end with a letter"),
			// validation.Match(regexp.MustCompile(`((([a-zA-Z])+([\s]?))+)`)).Error("invalid city name"),
		),

		validation.Field(&pharmaReq.SpecialLocationName,
			validation.Required,
			validation.Length(2, 100),
			validation.Match(regexp.MustCompile(`([A-Za-z\s])+`)).Error("must contain only letters"),
			// validation.Match(regexp.MustCompile("^[A-Z]+")).Error("name must start with an uppercase letter"),
			// validation.Match(regexp.MustCompile("[A-Za-z]+$")).Error("name must end with a letter"),
			// validation.Match(regexp.MustCompile(`((([a-zA-Z])+([\s]?))+)`)).Error("invalid city name"),
		),
	)
}

type Pharmacist struct {
	PharmacistID     uuid.UUID `json:"pharmacist_id"`
	PharmacyBranchID uuid.UUID `json:"pharmacy_branch_id"`
	Username         string    `json:"username"`
	Password         string    `json:"password,omitempty"`
	Email            string    `json:"email"`
	Role             string    `json:"role"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
}

type CreatePharmacistReq struct {
	PharmacyBranchName string `json:"pharmacy_branch_name"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	Role               string `json:"role"`
}

type CreatePharmacistRequest struct {
	PharmacyBranchID uuid.UUID `json:"pharmacy_branch_id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	Role             string    `json:"role"`
}

func (req CreatePharmacistReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username,
			validation.Required.Error("username is required"),
			validation.Length(5, 10).Error("username length must be between 5 and 10 characters"),
			is.Alphanumeric.Error("username must only contain letters and/or numbers")),
		validation.Field(&req.Password,
			validation.Required.Error("password is required"),
			validation.Length(8, 20).Error("password length must be between 8 and 10 characters"),
			validation.Match(regexp.MustCompile("([^A-Za-z0-9]+)")).Error("must have atleast one special character"),
			validation.Match(regexp.MustCompile("([A-Z]+)")).Error("must have atleast one uppercase letter"),
			validation.Match(regexp.MustCompile("([a-z]+)")).Error("must have atleast one lowercase letter"),
			validation.Match(regexp.MustCompile("([0-9]+)")).Error("must have atleast one digit")),
		validation.Field(&req.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("email must be valid")),
	)
}

type CreatePharmacyManagerRequest struct {
	PharmacBranchID uuid.UUID `json:"pharmacy_branch_id"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Email           string    `json:"email"`
	Role            string    `json:"role"`
}

func (usrReq CreatePharmacyManagerRequest) Validate() error {
	return validation.ValidateStruct(&usrReq,
		validation.Field(&usrReq.Username,
			validation.Required.Error("username is required"),
			validation.Length(5, 10).Error("must be between 5 and 10 characters long"),
			is.Alphanumeric.Error("must only contain numbers and/or letters")),
		validation.Field(&usrReq.Password,
			validation.Required.Error("password is required"),
			validation.Length(8, 20).Error("must be between 8 and 20 characters long"),
			validation.Match(regexp.MustCompile("([^A-Za-z0-9]+)")).Error("must have atleast one special character"),
			validation.Match(regexp.MustCompile("([A-Z]+)")).Error("must have atleast one uppercase letter"),
			validation.Match(regexp.MustCompile("([a-z]+)")).Error("must have atleast one lowercase letter"),
			validation.Match(regexp.MustCompile("([0-9]+)")).Error("must have atleast one digit")),
		validation.Field(&usrReq.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("must be a valid email")),
	)
}
