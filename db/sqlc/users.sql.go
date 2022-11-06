// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" (
  name,
  profile_text
) VALUES (
  $1, $2
)
RETURNING id, name, profile_text, profile_image, header_image, created_at
`

type CreateUserParams struct {
	Name        string `json:"name"`
	ProfileText string `json:"profileText"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.ProfileText)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ProfileText,
		&i.ProfileImage,
		&i.HeaderImage,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, profile_text, profile_image, header_image, created_at FROM "users"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ProfileText,
		&i.ProfileImage,
		&i.HeaderImage,
		&i.CreatedAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, profile_text, profile_image, header_image, created_at FROM "users"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]Users, error) {
	rows, err := q.db.QueryContext(ctx, listUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Users
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ProfileText,
			&i.ProfileImage,
			&i.HeaderImage,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserName = `-- name: UpdateUserName :one
UPDATE "users" 
set name = $2
WHERE id = $1
RETURNING id, name, profile_text, profile_image, header_image, created_at
`

type UpdateUserNameParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUserName, arg.ID, arg.Name)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ProfileText,
		&i.ProfileImage,
		&i.HeaderImage,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserProfile_text = `-- name: UpdateUserProfile_text :exec
UPDATE "users" 
set profile_text = $2
WHERE id = $1
`

type UpdateUserProfile_textParams struct {
	ID          int32  `json:"id"`
	ProfileText string `json:"profileText"`
}

func (q *Queries) UpdateUserProfile_text(ctx context.Context, arg UpdateUserProfile_textParams) error {
	_, err := q.db.ExecContext(ctx, updateUserProfile_text, arg.ID, arg.ProfileText)
	return err
}
