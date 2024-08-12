-- name: CheckAdmin :one
SELECT completed FROM setup WHERE id = 1;

-- name: MarkAdminAsCreated :exec
UPDATE setup SET completed = TRUE WHERE id = 1;
