// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: image_type.sql

package db

import (
	"context"
)

const getImageTypeById = `-- name: GetImageTypeById :one
SELECT id, name, description, variant FROM image_types WHERE id = $1 LIMIT 1
`

func (q *Queries) GetImageTypeById(ctx context.Context, id int32) (ImageType, error) {
	row := q.db.QueryRowContext(ctx, getImageTypeById, id)
	var i ImageType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Variant,
	)
	return i, err
}

const getImageTypeByName = `-- name: GetImageTypeByName :one
SELECT id, name, description, variant FROM image_types WHERE name = $1 LIMIT 1
`

func (q *Queries) GetImageTypeByName(ctx context.Context, name string) (ImageType, error) {
	row := q.db.QueryRowContext(ctx, getImageTypeByName, name)
	var i ImageType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Variant,
	)
	return i, err
}