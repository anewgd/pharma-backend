// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: pharmacist.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPharmacist = `-- name: CreatePharmacist :one
INSERT INTO pharmacists (
  pharmacy_branch_id,
  username,
  password,
  email,
  role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING pharmacist_id, pharmacy_branch_id, username, password, email, role, created_at
`

type CreatePharmacistParams struct {
	PharmacyBranchID uuid.UUID `json:"pharmacy_branch_id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	Role             string    `json:"role"`
}

func (q *Queries) CreatePharmacist(ctx context.Context, arg CreatePharmacistParams) (Pharmacist, error) {
	row := q.db.QueryRow(ctx, createPharmacist,
		arg.PharmacyBranchID,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.Role,
	)
	var i Pharmacist
	err := row.Scan(
		&i.PharmacistID,
		&i.PharmacyBranchID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getPharmacist = `-- name: GetPharmacist :one
SELECT pharmacist_id, pharmacy_branch_id, username, password, email, role, created_at FROM pharmacists
WHERE pharmacist_id = $1 LIMIT 1
`

func (q *Queries) GetPharmacist(ctx context.Context, pharmacistID uuid.UUID) (Pharmacist, error) {
	row := q.db.QueryRow(ctx, getPharmacist, pharmacistID)
	var i Pharmacist
	err := row.Scan(
		&i.PharmacistID,
		&i.PharmacyBranchID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}
