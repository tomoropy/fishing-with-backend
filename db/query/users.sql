-- name: CreateUser :one
INSERT INTO "users" (
  name,
  email,
  hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "users"
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM "users"
WHERE name = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM "users"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE "users"
SET
  profile_text = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE id = $1;
