-- name: CreatePhoto :one
INSERT INTO "photos" (
  user_id,
  image
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UserPhotoList :many
SELECT * FROM "photos"
WHERE user_id = $1
ORDER BY created_at
LIMIT $2;

-- name: DeletePhoto :exec
DELETE FROM "photos"
WHERE id = $1;
