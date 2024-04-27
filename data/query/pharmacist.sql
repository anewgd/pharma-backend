-- name: CreatePharmacist :one
INSERT INTO pharmacists (
  pharmacy_branch_id,
  username,
  password,
  email,
  role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPharmacist :one
SELECT * FROM pharmacists
WHERE pharmacist_id = $1 LIMIT 1;