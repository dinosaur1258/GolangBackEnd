// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: company.sql

package db

import (
	"context"
)

const delete = `-- name: Delete :exec
DELETE FROM company WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, delete, id)
	return err
}

const getAll = `-- name: GetAll :many
SELECT id, name, info FROM company
`

func (q *Queries) GetAll(ctx context.Context) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(&i.ID, &i.Name, &i.Info); err != nil {
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

const getByID = `-- name: GetByID :one
SELECT id, name, info FROM company WHERE id = $1
`

func (q *Queries) GetByID(ctx context.Context, id int32) (Company, error) {
	row := q.db.QueryRowContext(ctx, getByID, id)
	var i Company
	err := row.Scan(&i.ID, &i.Name, &i.Info)
	return i, err
}

const insert = `-- name: Insert :exec
INSERT INTO company (id, name, info) 
VALUES ($1, $2, $3)
`

type InsertParams struct {
	ID   int32
	Name string
	Info string
}

func (q *Queries) Insert(ctx context.Context, arg InsertParams) error {
	_, err := q.db.ExecContext(ctx, insert, arg.ID, arg.Name, arg.Info)
	return err
}

const update = `-- name: Update :execrows
Update company
SET
    name = $1,
    info = $2
WHERE
    id = $3
`

type UpdateParams struct {
	Name string
	Info string
	ID   int32
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, update, arg.Name, arg.Info, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
