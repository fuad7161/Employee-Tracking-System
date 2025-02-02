// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: sbu.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSBU = `-- name: CreateSBU :one
INSERT INTO sbus (
    sbu_name, sbu_head_user_id
) VALUES (
    $1, $2
) RETURNING id, sbu_name, sbu_head_user_id, created_at
`

type CreateSBUParams struct {
	SbuName       pgtype.Text `json:"sbu_name"`
	SbuHeadUserID pgtype.Int8 `json:"sbu_head_user_id"`
}

func (q *Queries) CreateSBU(ctx context.Context, arg CreateSBUParams) (Sbus, error) {
	row := q.db.QueryRow(ctx, createSBU, arg.SbuName, arg.SbuHeadUserID)
	var i Sbus
	err := row.Scan(
		&i.ID,
		&i.SbuName,
		&i.SbuHeadUserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSBU = `-- name: DeleteSBU :exec
DELETE FROM sbus WHERE id = $1
`

func (q *Queries) DeleteSBU(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSBU, id)
	return err
}

const getSBU = `-- name: GetSBU :one
SELECT id, sbu_name, sbu_head_user_id, created_at FROM sbus
WHERE id = $1
`

func (q *Queries) GetSBU(ctx context.Context, id int64) (Sbus, error) {
	row := q.db.QueryRow(ctx, getSBU, id)
	var i Sbus
	err := row.Scan(
		&i.ID,
		&i.SbuName,
		&i.SbuHeadUserID,
		&i.CreatedAt,
	)
	return i, err
}

const listSBUs = `-- name: ListSBUs :many
SELECT id, sbu_name, sbu_head_user_id, created_at FROM sbus
ORDER BY id
`

func (q *Queries) ListSBUs(ctx context.Context) ([]Sbus, error) {
	rows, err := q.db.Query(ctx, listSBUs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Sbus{}
	for rows.Next() {
		var i Sbus
		if err := rows.Scan(
			&i.ID,
			&i.SbuName,
			&i.SbuHeadUserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSBUHeadUserID = `-- name: UpdateSBUHeadUserID :exec
UPDATE sbus SET sbu_head_user_id = $2
WHERE id = $1
`

type UpdateSBUHeadUserIDParams struct {
	ID            int64       `json:"id"`
	SbuHeadUserID pgtype.Int8 `json:"sbu_head_user_id"`
}

func (q *Queries) UpdateSBUHeadUserID(ctx context.Context, arg UpdateSBUHeadUserIDParams) error {
	_, err := q.db.Exec(ctx, updateSBUHeadUserID, arg.ID, arg.SbuHeadUserID)
	return err
}

const updateSbuByID = `-- name: UpdateSbuByID :exec
UPDATE sbus SET sbu_name =$2, sbu_head_user_id = $3
WHERE id = $1
`

type UpdateSbuByIDParams struct {
	ID            int64       `json:"id"`
	SbuName       pgtype.Text `json:"sbu_name"`
	SbuHeadUserID pgtype.Int8 `json:"sbu_head_user_id"`
}

func (q *Queries) UpdateSbuByID(ctx context.Context, arg UpdateSbuByIDParams) error {
	_, err := q.db.Exec(ctx, updateSbuByID, arg.ID, arg.SbuName, arg.SbuHeadUserID)
	return err
}
