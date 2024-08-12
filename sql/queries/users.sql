-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, email, api_key, usertype)
VALUES ($1, $2, $3, $4, $5,
    encode(sha256(random()::text::bytea), 'hex'),
    $6
)
RETURNING *;

-- name: GerUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;
