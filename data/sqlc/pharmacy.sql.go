// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: pharmacy.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPharmacy = `-- name: CreatePharmacy :one
INSERT INTO pharmacies (
  pharmacy_name,
  user_id
) VALUES (
  $1, $2
) RETURNING pharmacy_id, pharmacy_name, user_id, created_at
`

type CreatePharmacyParams struct {
	PharmacyName string    `json:"pharmacy_name"`
	UserID       uuid.UUID `json:"user_id"`
}

func (q *Queries) CreatePharmacy(ctx context.Context, arg CreatePharmacyParams) (Pharmacy, error) {
	row := q.db.QueryRow(ctx, createPharmacy, arg.PharmacyName, arg.UserID)
	var i Pharmacy
	err := row.Scan(
		&i.PharmacyID,
		&i.PharmacyName,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const createPharmacyBranch = `-- name: CreatePharmacyBranch :one
INSERT INTO pharmacy_branches (
  pharmacy_id,
  pharmacy_branch_name,
  city,
  sub_city,
  special_location_name
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING pharmacy_branch_id, pharmacy_id, pharmacy_branch_name, city, sub_city, special_location_name, created_at
`

type CreatePharmacyBranchParams struct {
	PharmacyID          uuid.UUID `json:"pharmacy_id"`
	PharmacyBranchName  string    `json:"pharmacy_branch_name"`
	City                string    `json:"city"`
	SubCity             string    `json:"sub_city"`
	SpecialLocationName string    `json:"special_location_name"`
}

func (q *Queries) CreatePharmacyBranch(ctx context.Context, arg CreatePharmacyBranchParams) (PharmacyBranch, error) {
	row := q.db.QueryRow(ctx, createPharmacyBranch,
		arg.PharmacyID,
		arg.PharmacyBranchName,
		arg.City,
		arg.SubCity,
		arg.SpecialLocationName,
	)
	var i PharmacyBranch
	err := row.Scan(
		&i.PharmacyBranchID,
		&i.PharmacyID,
		&i.PharmacyBranchName,
		&i.City,
		&i.SubCity,
		&i.SpecialLocationName,
		&i.CreatedAt,
	)
	return i, err
}

const getPharmacy = `-- name: GetPharmacy :one
SELECT pharmacy_id, pharmacy_name, user_id, created_at FROM pharmacies
WHERE pharmacy_id = $1 LIMIT 1
`

func (q *Queries) GetPharmacy(ctx context.Context, pharmacyID uuid.UUID) (Pharmacy, error) {
	row := q.db.QueryRow(ctx, getPharmacy, pharmacyID)
	var i Pharmacy
	err := row.Scan(
		&i.PharmacyID,
		&i.PharmacyName,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getPharmacyBrachByName = `-- name: GetPharmacyBrachByName :one
SELECT pharmacy_branch_id, pharmacy_id, pharmacy_branch_name, city, sub_city, special_location_name, created_at FROM pharmacy_branches
WHERE pharmacy_branch_name = $1 LIMIT 1
`

func (q *Queries) GetPharmacyBrachByName(ctx context.Context, pharmacyBranchName string) (PharmacyBranch, error) {
	row := q.db.QueryRow(ctx, getPharmacyBrachByName, pharmacyBranchName)
	var i PharmacyBranch
	err := row.Scan(
		&i.PharmacyBranchID,
		&i.PharmacyID,
		&i.PharmacyBranchName,
		&i.City,
		&i.SubCity,
		&i.SpecialLocationName,
		&i.CreatedAt,
	)
	return i, err
}

const getPharmacyByAdminID = `-- name: GetPharmacyByAdminID :one
SELECT pharmacy_id, pharmacy_name, user_id, created_at FROM pharmacies
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetPharmacyByAdminID(ctx context.Context, userID uuid.UUID) (Pharmacy, error) {
	row := q.db.QueryRow(ctx, getPharmacyByAdminID, userID)
	var i Pharmacy
	err := row.Scan(
		&i.PharmacyID,
		&i.PharmacyName,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
