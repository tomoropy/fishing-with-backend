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
  profile_text,
  email,
  hashed_password
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, profile_text, profile_image, header_image, created_at, email, hashed_password
`

type CreateUserParams struct {
	Name           string `json:"name"`
	ProfileText    string `json:"profileText"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.ProfileText,
		arg.Email,
		arg.HashedPassword,
	)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ProfileText,
		&i.ProfileImage,
		&i.HeaderImage,
		&i.CreatedAt,
		&i.Email,
		&i.HashedPassword,
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
SELECT id, name, profile_text, profile_image, header_image, created_at, email, hashed_password FROM "users"
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
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, name, profile_text, profile_image, header_image, created_at, email, hashed_password FROM "users"
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, name)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ProfileText,
		&i.ProfileImage,
		&i.HeaderImage,
		&i.CreatedAt,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, profile_text, profile_image, header_image, created_at, email, hashed_password FROM "users"
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
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ProfileText,
			&i.ProfileImage,
			&i.HeaderImage,
			&i.CreatedAt,
			&i.Email,
			&i.HashedPassword,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE "users"
SET
  name = $2,
  profile_text = $3,
  email = $4,
  hashed_password = $5
WHERE id = $1
`

type UpdateUserParams struct {
	ID             int32  `json:"id"`
	Name           string `json:"name"`
	ProfileText    string `json:"profileText"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.ProfileText,
		arg.Email,
		arg.HashedPassword,
	)
	return err
}
