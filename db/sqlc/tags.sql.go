// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: tags.sql

package db

import (
	"context"
	"database/sql"
)

const createModuleTag = `-- name: CreateModuleTag :one
INSERT INTO module_tags (module_id, tag_id)
VALUES ($1, $2)
RETURNING module_id, tag_id
`

type CreateModuleTagParams struct {
	ModuleID int32 `json:"module_id"`
	TagID    int32 `json:"tag_id"`
}

func (q *Queries) CreateModuleTag(ctx context.Context, arg CreateModuleTagParams) (ModuleTag, error) {
	row := q.db.QueryRowContext(ctx, createModuleTag, arg.ModuleID, arg.TagID)
	var i ModuleTag
	err := row.Scan(&i.ModuleID, &i.TagID)
	return i, err
}

const createModuleTypeTagAvailable = `-- name: CreateModuleTypeTagAvailable :one
INSERT INTO module_type_tags_available (module_type, tag)
VALUES ($1, $2)
RETURNING id, module_type, tag
`

type CreateModuleTypeTagAvailableParams struct {
	ModuleType ModuleType `json:"module_type"`
	Tag        string     `json:"tag"`
}

func (q *Queries) CreateModuleTypeTagAvailable(ctx context.Context, arg CreateModuleTypeTagAvailableParams) (ModuleTypeTagsAvailable, error) {
	row := q.db.QueryRowContext(ctx, createModuleTypeTagAvailable, arg.ModuleType, arg.Tag)
	var i ModuleTypeTagsAvailable
	err := row.Scan(&i.ID, &i.ModuleType, &i.Tag)
	return i, err
}

const deleteModuleTag = `-- name: DeleteModuleTag :exec
DELETE FROM module_tags
WHERE
        module_id = COALESCE($1, module_id) AND
        tag_id = COALESCE($2, tag_id) AND
    (NOT ($1 IS NULL AND $2 IS NULL))
`

type DeleteModuleTagParams struct {
	ModuleID sql.NullInt32 `json:"module_id"`
	TagID    sql.NullInt32 `json:"tag_id"`
}

func (q *Queries) DeleteModuleTag(ctx context.Context, arg DeleteModuleTagParams) error {
	_, err := q.db.ExecContext(ctx, deleteModuleTag, arg.ModuleID, arg.TagID)
	return err
}

const deleteModuleTypeTagAvailable = `-- name: DeleteModuleTypeTagAvailable :exec
DELETE FROM module_type_tags_available WHERE id = $1
`

func (q *Queries) DeleteModuleTypeTagAvailable(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteModuleTypeTagAvailable, id)
	return err
}

const getModuleTypeTagAvailable = `-- name: GetModuleTypeTagAvailable :one
SELECT id, module_type, tag, count FROM view_module_type_tags_available WHERE id = $1
`

func (q *Queries) GetModuleTypeTagAvailable(ctx context.Context, tagID int32) (ViewModuleTypeTagsAvailable, error) {
	row := q.db.QueryRowContext(ctx, getModuleTypeTagAvailable, tagID)
	var i ViewModuleTypeTagsAvailable
	err := row.Scan(
		&i.ID,
		&i.ModuleType,
		&i.Tag,
		&i.Count,
	)
	return i, err
}

const getModuleTypeTagsAvailable = `-- name: GetModuleTypeTagsAvailable :many
SELECT id, module_type, tag, count FROM view_module_type_tags_available
WHERE module_type = $1
`

func (q *Queries) GetModuleTypeTagsAvailable(ctx context.Context, moduleType ModuleType) ([]ViewModuleTypeTagsAvailable, error) {
	rows, err := q.db.QueryContext(ctx, getModuleTypeTagsAvailable, moduleType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ViewModuleTypeTagsAvailable{}
	for rows.Next() {
		var i ViewModuleTypeTagsAvailable
		if err := rows.Scan(
			&i.ID,
			&i.ModuleType,
			&i.Tag,
			&i.Count,
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

const updateModuleTypeTagAvailable = `-- name: UpdateModuleTypeTagAvailable :one
UPDATE module_type_tags_available
SET tag = $1
WHERE id = $2
RETURNING id, module_type, tag
`

type UpdateModuleTypeTagAvailableParams struct {
	Tag string `json:"tag"`
	ID  int32  `json:"id"`
}

func (q *Queries) UpdateModuleTypeTagAvailable(ctx context.Context, arg UpdateModuleTypeTagAvailableParams) (ModuleTypeTagsAvailable, error) {
	row := q.db.QueryRowContext(ctx, updateModuleTypeTagAvailable, arg.Tag, arg.ID)
	var i ModuleTypeTagsAvailable
	err := row.Scan(&i.ID, &i.ModuleType, &i.Tag)
	return i, err
}