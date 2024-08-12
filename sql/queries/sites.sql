-- name: CreateSite :one
INSERT INTO sites (id, created_at, updated_at, name, url, added_by_user)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetSites :many
SELECT * from sites;

-- name: GetMyAddedSites :many
SELECT * from sites where added_by_user = $1;

-- name: GetAllSites :many
SELECT * from sites;