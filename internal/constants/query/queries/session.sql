-- name: CreateUserSession :one
INSERT INTO user_sessions (
    user_id,
    refresh_token,
    expires_at
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserSession :one
SELECT * FROM user_sessions
WHERE session_id = $1 LIMIT 1;

-- name: DeleteUserSession :exec
DELETE FROM user_sessions
WHERE user_id = $1;

-- name: CreatePharmacistSession :one
INSERT INTO pharmacist_sessions (
    session_id,
    pharmacist_id,
    refresh_token,
    expires_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPharmacistSession :one
SELECT * FROM pharmacist_sessions
WHERE session_id = $1 LIMIT 1;

-- name: DeletePharmacistSession :exec
DELETE FROM pharmacist_sessions
WHERE pharmacist_id = $1;

-- name: DeleteAllUserSessions :exec
TRUNCATE user_sessions;

-- name: DeleteAllPharmacistSessions :exec
TRUNCATE pharmacist_sessions;
