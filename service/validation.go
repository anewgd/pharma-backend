package service

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (pharmaReq CreatePharmacyRequest) Validate() error {
	return validation.ValidateStruct(&pharmaReq,
		validation.Field(&pharmaReq.PharmacyName,
			validation.Required,
			validation.Length(5, 30),
			validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
			validation.Match(regexp.MustCompile("([A-Za-z])+")).Error("must contain atleast one or more letters"),
		),
	)
}

func (pharmaReq CreatePharmacyBranchRequest) Validate() error {
	return validation.ValidateStruct(&pharmaReq,
		validation.Field(&pharmaReq.PharmacyBranchName,
			validation.Required,
			validation.Length(5, 30),
			validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
			validation.Match(regexp.MustCompile("([A-Za-z])+")).Error("must contain atleast one or more letters"),
		),
		validation.Field(&pharmaReq.City,
			validation.Required,
			validation.Length(2, 40),
			validation.Match(regexp.MustCompile("^[A-Z]+")).Error("name must start with an uppercase letter"),
			validation.Match(regexp.MustCompile("[A-Za-z]+$")).Error("name must end with a letter"),
			validation.Match(regexp.MustCompile(`((([a-zA-Z])+([\s]?))+)`)).Error("invalid city name"),
		),
		validation.Field(&pharmaReq.SubCity,
			validation.Required,
			validation.Length(2, 40),
			validation.Match(regexp.MustCompile("^[A-Z]+")).Error("name must start with an uppercase letter"),
			validation.Match(regexp.MustCompile("[A-Za-z]+$")).Error("name must end with a letter"),
			validation.Match(regexp.MustCompile(`((([a-zA-Z])+([\s]?))+)`)).Error("invalid city name"),
		),

		validation.Field(&pharmaReq.SpecialLocationName,
			validation.Required,
			validation.Length(2, 40),
			validation.Match(regexp.MustCompile("^[A-Z]+")).Error("name must start with an uppercase letter"),
			validation.Match(regexp.MustCompile("[A-Za-z]+$")).Error("name must end with a letter"),
			validation.Match(regexp.MustCompile(`((([a-zA-Z])+([\s]?))+)`)).Error("invalid city name"),
		),
	)
}

func (c CreateDrugRequest) Validate() error {
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
			validation.Max(100000),
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

func (usrReq CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&usrReq,
		validation.Field(&usrReq.Username, validation.Required, validation.Length(5, 10), is.Alphanumeric),
		validation.Field(&usrReq.Password,
			validation.Required,
			validation.Length(8, 15),
			validation.Match(regexp.MustCompile("([^A-Za-z0-9]+)")).Error("must have atleast one special character"),
			validation.Match(regexp.MustCompile("([A-Z]+)")).Error("must have atleast one uppercase letter"),
			validation.Match(regexp.MustCompile("([a-z]+)")).Error("must have atleast one lowercase letter"),
			validation.Match(regexp.MustCompile("([0-9]+)")).Error("must have atleast one digit")),
		validation.Field(&usrReq.Email, validation.Required, is.Email),
	)
}

func (usrReq CreatePharmacyManagerRequest) Validate() error {
	return validation.ValidateStruct(&usrReq,
		validation.Field(&usrReq.Username, validation.Required, validation.Length(5, 10), is.Alphanumeric),
		validation.Field(&usrReq.Password,
			validation.Required,
			validation.Length(8, 15),
			validation.Match(regexp.MustCompile("([^A-Za-z0-9]+)")).Error("must have atleast one special character"),
			validation.Match(regexp.MustCompile("([A-Z]+)")).Error("must have atleast one uppercase letter"),
			validation.Match(regexp.MustCompile("([a-z]+)")).Error("must have atleast one lowercase letter"),
			validation.Match(regexp.MustCompile("([0-9]+)")).Error("must have atleast one digit")),
		validation.Field(&usrReq.Email, validation.Required, is.Email),
		validation.Field(&usrReq.PharmacyBranchName,
			validation.Required,
			validation.Length(5, 30),
			validation.Match(regexp.MustCompile(`([A-Za-z0-9\s])+`)).Error("must contain only letters and/or numbers"),
			validation.Match(regexp.MustCompile("([A-Za-z])+")).Error("must contain atleast one or more letters"),
		),
	)
}

func (req LoginUserRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, validation.Length(5, 10), is.Alphanumeric),
		validation.Field(&req.Password,
			validation.Required,
			validation.Length(8, 15)),
	)
}
