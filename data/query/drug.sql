-- name: CreateDrug :one
INSERT INTO drugs (
  pharmacy_branch_id,
  brand_name,
  generic_name,
  quantity,
  expiration_date,
  manufacturing_date,
  pharmacist_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;