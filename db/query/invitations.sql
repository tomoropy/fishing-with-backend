-- name: CreateInvitation :one
INSERT INTO "invitations" (
  title,
  place,
  people,
  start_time,
  end_time,
  transpost,
  comment,
  applicant,
  image
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetInvitation :one
SELECT * FROM "invitations"
WHERE id = $1 LIMIT 1;

-- name: ListInvitation :many
SELECT * FROM "invitations"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListInvitationByUser :many
SELECT * FROM "invitations"
WHERE user_id = $1
ORDER BY id
LIMIT $2;

-- name: UpdateInvitation :exec
UPDATE "invitations" 
set   
  title = $2,
  place = $3,
  people = $4,
  start_time = $5,
  end_time = $6,
  transpost = $7,
  comment = $8,
  applicant = $9,
  image = $10
WHERE id = $1
RETURNING *;

-- name: DeleteInvitation :exec
DELETE FROM "invitations"
WHERE id = $1;
