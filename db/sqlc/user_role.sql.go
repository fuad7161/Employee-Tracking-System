// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_role.sql

package db

import (
	"context"
)

const createUserRole = `-- name: CreateUserRole :one
INSERT INTO user_roles (
    role_name
) VALUES (
    $1
) RETURNING id, role_name, created_at
`

func (q *Queries) CreateUserRole(ctx context.Context, roleName string) (UserRole, error) {
	row := q.db.QueryRow(ctx, createUserRole, roleName)
	var i UserRole
	err := row.Scan(&i.ID, &i.RoleName, &i.CreatedAt)
	return i, err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM user_roles WHERE id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteRole, id)
	return err
}

const getUserRoleByID = `-- name: GetUserRoleByID :one
SELECT id, role_name, created_at FROM user_roles
WHERE id = $1
`

func (q *Queries) GetUserRoleByID(ctx context.Context, id int64) (UserRole, error) {
	row := q.db.QueryRow(ctx, getUserRoleByID, id)
	var i UserRole
	err := row.Scan(&i.ID, &i.RoleName, &i.CreatedAt)
	return i, err
}

const listUserRole = `-- name: ListUserRole :many
SELECT id, role_name, created_at FROM user_roles
ORDER BY id
`

func (q *Queries) ListUserRole(ctx context.Context) ([]UserRole, error) {
	rows, err := q.db.Query(ctx, listUserRole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserRole{}
	for rows.Next() {
		var i UserRole
		if err := rows.Scan(&i.ID, &i.RoleName, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserByRole = `-- name: UpdateUserByRole :exec
UPDATE user_roles SET role_name = $2
WHERE id = $1
`

type UpdateUserByRoleParams struct {
	ID       int64  `json:"id"`
	RoleName string `json:"role_name"`
}

func (q *Queries) UpdateUserByRole(ctx context.Context, arg UpdateUserByRoleParams) error {
	_, err := q.db.Exec(ctx, updateUserByRole, arg.ID, arg.RoleName)
	return err
}
