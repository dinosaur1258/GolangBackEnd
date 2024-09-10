-- name: GetAll :many
SELECT * FROM company;

-- name: GetByID :one
SELECT id, name, info FROM company WHERE id = $1;

-- name: Delete :exec
DELETE FROM company WHERE id = $1;

-- name: Insert :exec
INSERT INTO company (id, name, info) 
VALUES ($1, $2, $3);

-- name: Update :execrows
Update company
SET
    name = $1,
    info = $2
WHERE
    id = $3;