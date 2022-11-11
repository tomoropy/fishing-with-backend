-- name: CreateUser :one
INSERT INTO "users" (
  name,
  profile_text,
  email,
  hashed_password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "users"
WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM "users"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserName :one
UPDATE "users" 
set name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserProfile_text :exec
UPDATE "users" 
set profile_text = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE id = $1;
