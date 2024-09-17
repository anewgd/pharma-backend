package dto

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type CreateUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (usrReq CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&usrReq,
		validation.Field(&usrReq.Username,
			validation.Required.Error("username is required"),
			validation.Length(5, 10).Error("username length must be between 5 and 10 characters"),
			is.Alphanumeric.Error("username must only contain letters and/or numbers")),
		validation.Field(&usrReq.Password,
			validation.Required.Error("password is required"),
			validation.Length(8, 20).Error("password length must be between 8 and 10 characters"),
			validation.Match(regexp.MustCompile("([^A-Za-z0-9]+)")).Error("must have atleast one special character"),
			validation.Match(regexp.MustCompile("([A-Z]+)")).Error("must have atleast one uppercase letter"),
			validation.Match(regexp.MustCompile("([a-z]+)")).Error("must have atleast one lowercase letter"),
			validation.Match(regexp.MustCompile("([0-9]+)")).Error("must have atleast one digit")),
		validation.Field(&usrReq.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("email must be valid")),
	)
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u LoginUserRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required.Error("username is required")),
		validation.Field(&u.Password, validation.Required.Error("password is required")),
	)
}

type LoginUserResponse struct {
	Username              string    `json:"username"`
	Email                 string    `json:"email"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}
