-- name: CreatePharmacy :one
INSERT INTO pharmacies (
  pharmacy_name,
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetPharmacy :one
SELECT * FROM pharmacies
WHERE pharmacy_id = $1 LIMIT 1;

-- name: GetPharmacyByAdminID :one
SELECT * FROM pharmacies
WHERE user_id = $1 LIMIT 1;

-- name: CreatePharmacyBranch :one
INSERT INTO pharmacy_branches (
  pharmacy_id,
  pharmacy_branch_name,
  city,
  sub_city,
  special_location_name
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPharmacyBrachByName :one
SELECT * FROM pharmacy_branches
WHERE pharmacy_branch_name = $1 LIMIT 1;

