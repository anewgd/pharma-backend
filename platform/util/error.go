package util

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joomcode/errorx"
)

type ErrorResponse struct {
	Err        error
	StatusCode int
	Msg        string
}

var AuthenticationError = errorx.NewType(errorx.CommonErrors, "authentication_error", errorx.CaseNoTrait())
var AuthorizationError = errorx.NewType(errorx.CommonErrors, "authorization_error", errorx.CaseNoTrait())
var RequestError = errorx.NewType(errorx.CommonErrors, "request_error", errorx.CaseNoTrait())

func (e ErrorResponse) Error() string {
	return e.Msg
}

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}

func NewErrorResponse(err error, statusCode int, msg string) ErrorResponse {
	return ErrorResponse{
		Err:        err,
		StatusCode: statusCode,
		Msg:        msg,
	}
}
