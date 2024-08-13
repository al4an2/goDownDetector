-- name: CreateSiteFollow :one
INSERT INTO site_follows(
    id,
    created_at,
    updated_at,
    user_id,
    site_id)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSiteFollows :many
SELECT * FROM site_follows WHERE user_id = $1;

-- name: GetAllSiteFollows :many
SELECT * FROM site_follows;

-- name: DeleteSiteFollow :exec
DELETE FROM site_follows 
USING users 
WHERE site_follows.user_id = users.user_id
AND site_follows.id = $1 
AND (site_follows.user_id = $2 or users.usertype = "admin");