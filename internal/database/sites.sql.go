// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sites.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSite = `-- name: CreateSite :one
INSERT INTO sites (id, created_at, updated_at, name, url, added_by_user)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, name, url, added_by_user
`

type CreateSiteParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Url         string
	AddedByUser uuid.UUID
}

func (q *Queries) CreateSite(ctx context.Context, arg CreateSiteParams) (Site, error) {
	row := q.db.QueryRowContext(ctx, createSite,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.AddedByUser,
	)
	var i Site
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.AddedByUser,
	)
	return i, err
}

const getAllSites = `-- name: GetAllSites :many
SELECT id, created_at, updated_at, name, url, added_by_user from sites
`

func (q *Queries) GetAllSites(ctx context.Context) ([]Site, error) {
	rows, err := q.db.QueryContext(ctx, getAllSites)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Site
	for rows.Next() {
		var i Site
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.AddedByUser,
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

const getMyAddedSites = `-- name: GetMyAddedSites :many
SELECT id, created_at, updated_at, name, url, added_by_user from sites where added_by_user = $1
`

func (q *Queries) GetMyAddedSites(ctx context.Context, addedByUser uuid.UUID) ([]Site, error) {
	rows, err := q.db.QueryContext(ctx, getMyAddedSites, addedByUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Site
	for rows.Next() {
		var i Site
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.AddedByUser,
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

const getSites = `-- name: GetSites :many
SELECT id, created_at, updated_at, name, url, added_by_user from sites
`

func (q *Queries) GetSites(ctx context.Context) ([]Site, error) {
	rows, err := q.db.QueryContext(ctx, getSites)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Site
	for rows.Next() {
		var i Site
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.AddedByUser,
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
