package errors

import (
	"errors"
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/joomcode/errorx"
)

var ErrorMap = map[*errorx.Type]int{
	ErrInvalidUserInput: http.StatusBadRequest,
	ErrDataExists:       http.StatusBadRequest,
	ErrReadError:        http.StatusInternalServerError,
	ErrWriteError:       http.StatusInternalServerError,
	ErrNoRecordFound:    http.StatusNotFound,
	ErrUnauthorized:     http.StatusUnauthorized,
	ErrInvalidToken:     http.StatusUnauthorized,
	ErrAuthError:        http.StatusUnauthorized,
}

var (
	invalidInput = errorx.NewNamespace("validation error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	dbError      = errorx.NewNamespace("db error")
	duplicate    = errorx.NewNamespace("duplicate").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	dataNotFound = errorx.NewNamespace("data not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	authError    = errorx.NewNamespace("auth error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
)

var (
	ErrInvalidUserInput = errorx.NewType(invalidInput, "invalid user input")
	ErrWriteError       = errorx.NewType(dbError, "could not write to db")
	ErrReadError        = errorx.NewType(dbError, "could not read data from db")
	ErrDataExists       = errorx.NewType(duplicate, "data already exists")
	ErrNoRecordFound    = errorx.NewType(dataNotFound, "no record found")
	ErrUnauthorized     = errorx.NewType(authError, "could not authorize")
	ErrInvalidToken     = errorx.NewType(authError, "invalid token")
	ErrAuthError        = errorx.NewType(authError, "you are not authorized.")
)

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
